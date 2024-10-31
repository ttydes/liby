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

	db "github.com/ttydes/liby/services/account/db"
	pb "github.com/ttydes/liby/services/account/protogen"
)

type server struct {
	pb.UnimplementedAccountServiceServer
	db *pgxpool.Pool
}

var port = flag.Int("port", 8000, "the port to serve on")

func (s *server) MkAccount(ctx context.Context, req *pb.MkAccountReq) (*pb.MkAccountResp, error) {
	return db.MkAccount(ctx, s.db, req)
}

func (s *server) UpdateAccount(ctx context.Context, req *pb.UpdateAccountReq) (*pb.UpdateAccountResp, error) {
	return db.UpdateAccount(ctx, s.db, req)
}

func (s *server) DeleteAccount(ctx context.Context, req *pb.DeleteAccountReq) (*pb.DeleteAccountResp, error) {
	return db.DeleteAccount(ctx, s.db, req)
}

func (s *server) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.GetAccountResp, error) {
	return db.GetAccount(ctx, s.db, req)
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
	pb.RegisterAccountServiceServer(srv, &server{db: dbConn})

	// Start the server on the port defined earlier.
	log.Printf("Starting gRPC server on port %d...", *port)
	if err := srv.Serve(sock); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
