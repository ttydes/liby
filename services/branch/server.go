package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	db "github.com/ttydes/liby/services/branch/db"
	pb "github.com/ttydes/liby/services/branch/protogen"
)

type server struct {
	pb.UnimplementedBranchServiceServer
	db *pgxpool.Pool
}

var port = flag.Int("port", 8002, "the port to serve on")

func (s *server) MkBranch(ctx context.Context, req *pb.MkBranchReq) (*pb.MkBranchResp, error) {
	return db.MkBranch(ctx, s.db, req)
}

func (s *server) GetBranch(ctx context.Context, req *pb.GetBranchReq) (*pb.GetBranchResp, error) {
	return db.GetBranch(ctx, s.db, req)
}

func (s *server) UpdateBranch(ctx context.Context, req *pb.UpdateBranchReq) (*pb.UpdateBranchResp, error) {
	return db.UpdateBranch(ctx, s.db, req)
}

func (s *server) DeleteBranch(ctx context.Context, req *pb.DeleteBranchReq) (*pb.DeleteBranchResp, error) {
	return db.DeleteBranch(ctx, s.db, req)
}

func main() {
	ctx := context.Background()
	flag.Parse()

	// Make connection to db
	dbConn, err := db.MkClient(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Open a tcp socket
	sock, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Load TLS certificates so browsers don't complain about insecure site.
	creds, err := credentials.NewServerTLSFromFile("../../certs/server-cert.pem", "../../certs/server-key.pem")
	if err != nil {
		log.Fatalf("Failed to create credentials: %v", err)
	}

	// Initialize the server struct with the database connection
	srv := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterBranchServiceServer(srv, &server{db: dbConn})

	// Start the server on the port defined earlier.
	log.Printf("Starting gRPC server on port %d...", *port)
	if err := srv.Serve(sock); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
