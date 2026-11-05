package main

import (
	"auth-service/internal/config"
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/internal/token"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.Seller{}); err != nil {
		log.Fatalf("Failed to auto-migrate model: %v", err)
	}
	log.Println("Database connection established and model auto-migrated")

	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 5 {
		log.Fatal("JWT_SECRET must be set and at least 5 characters")
	}

	tokenManager := token.NewManager(secret, 24*time.Hour)

	userRepo := repository.NewUserRepository(db)
	sellerRepo := repository.NewSellerRepository(db)

	authService := service.NewAuthService(userRepo, tokenManager)
	sellerService := service.NewSellerService(sellerRepo, tokenManager)

	server := rpc.NewServer()
	if err := server.RegisterName("AuthService", authService); err != nil {
		log.Fatalf("Failed to register rpc auth service: %v", err)
	}

	if err := server.RegisterName("SellerService", sellerService); err != nil {
		log.Fatalf("Failed to register rpc seller service: %v", err)
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
