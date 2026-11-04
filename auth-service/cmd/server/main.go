package main

import (
	"auth-service/internal/config"
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to auto-migrate model: %v", err)
	}
	log.Println("Database connection established and model auto-migrated")

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)

	server := rpc.NewServer()
	if err := server.RegisterName("AuthService", authService); err != nil {
		log.Fatalf("Failed to register RPC service: %v", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Println("Auth RPC server listening on :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go server.ServeConn(conn)
	}
}
