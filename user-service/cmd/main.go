package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ALIK/user-service/internal/user"
	"github.com/ALIK/user-service/internal/grpc"
)

func main() {
	// Initialize logger
	logger := log.New(os.Stdout, "[user-service] ", log.LstdFlags|log.Lshortfile)

	// Initialize user service
	userService := user.NewUserService(logger)

	// Start gRPC server
	go func() {
		grpcServer := grpc.NewGRPCServer(userService, logger)
		err := grpcServer.Serve()
		if err != nil {

