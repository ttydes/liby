package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	bookpb "github.com/ttydes/liby/protogen/go/bookpb"
)

type Book struct {
	bookpb.UnimplementedBookServiceServer
	// books map[string]*book.Book
}

func (b *Book) CreateBook(ctx context.Context, req *bookpb.CreateBookReq) (*bookpb.CreateBookResp, error) {
	book := req.GetBook()
	if book.Id == "" {
		book.Id = generateID() // Implement generateID() to generate a unique ID
	}
	b.books[book.Id] = book
	return &book.CreateBookResp{Id: book.Id}, nil
}

func (b *Book) GetBook(ctx context.Context, req *book.GetBookReq) (*book.GetBookResp, error) {
	book, exists := b.books[req.GetId()]
	if !exists {
		return nil, grpc.Errorf(grpc.Code(grpc.NotFound), "Book not found")
	}
	return &book.GetBookResp{Book: book}, nil
}

func (b *Book) UpdateBook(ctx context.Context, req *book.UpdateBookReq) (*book.UpdateBookResp, error) {
	book := req.GetBook()
	if _, exists := b.books[book.Id]; !exists {
		return nil, grpc.Errorf(grpc.Code(grpc.NotFound), "Book not found")
	}
	b.books[book.Id] = book
	return &book.UpdateBookResp{Success: true}, nil
}

func (b *Book) DeleteBook(ctx context.Context, req *pb_book.DeleteBookReq) (*pb_book.DeleteBookResp, error) {
	id := req.GetId()
	if _, exists := b.books[id]; !exists {
		return nil, grpc.Errorf(grpc.Code(grpc.NotFound), "Book not found")
	}
	delete(b.books, id)
	return &pb_book.DeleteBookResp{Success: true}, nil
}

func (b *Book) ListBooks(ctx context.Context, req *pb_book.ListBooksReq) (*pb_book.ListBooksResp, error) {
	var books []*pb_book.Book
	for _, book := range b.books {
		books = append(books, book)
	}
	return &pb_book.ListBooksResp{Books: books}, nil
}

func generateID() string {
	return "unique-id"
}

func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb_book.RegisterBookServiceServer(s, &Book{books: make(map[string]*pb_book.Book)})
	reflection.Register(s)
	log.Println("Starting gRPC server on port 7777...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
