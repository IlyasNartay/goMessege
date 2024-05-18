package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/to/project/chat-service/internal/chat"
	"path/to/project/chat-service/internal/grpc"
	pb "path/to/project/chat-service/internal/grpc"

	"google.golang.org/grpc"
)

func main() {
	logger := log.New(os.Stdout, "chat-service ", log.LstdFlags)
	chatService := chat.NewChatService(logger)
	chatHandler := chat.NewChatHandler(chatService, logger)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			logger.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		pb.RegisterChatServiceServer(s, grpc.NewChatServer(chatService))

		logger.Println("gRPC server is running on port :50051...")
		if err := s.Serve(lis); err != nil {
			logger.Fatalf("failed to serve: %v", err)
		}
	}()

	http.HandleFunc("/messages", chatHandler.SendMessage)
	http.HandleFunc("/messages", chatHandler.GetMessages)

	logger.Println("HTTP server is running on port :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
