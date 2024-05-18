package main

import (
	"database/sql"
	"log"
	"net"

	pb "awesomeProject/auth_service/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db, err := sql.Open("postgres", "postgresql://postgres:9999@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	defer db.Close()

	// Проверьте соединение с базой данных
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, pb.Server{Db: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
