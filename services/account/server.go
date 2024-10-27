package main

import (
	db "github.com/ttydes/liby/services/account/db"
	pb "github.com/ttydes/liby/services/account/protogen"

	"context"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAccountServiceServer
	db *sqlx.DB
}

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

	// make connection to db
	db, err := db.MkClient(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// open a tcp socket
	sock, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv, &server{})

	log.Println("Starting gRPC server on port 8888...")
	if err := srv.Serve(sock); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
