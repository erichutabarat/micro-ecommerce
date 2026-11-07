package main

import (
	"log"
	"net"
	"net/rpc"
	"os"
	"product-service/internal/config"
	"product-service/internal/model"
	"product-service/internal/repository"
	"product-service/internal/service"
)

func main() {
	// init db
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalf("Failed to auto-migrate model: %v", err)
	}
	log.Println("Database connection established and model auto-migrated")

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)

	server := rpc.NewServer()
	if err := server.RegisterName("ProductService", productService); err != nil {
		log.Fatalf("Failed to register ProductService: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1235"
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Printf("Product RPC server listening on :%s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go server.ServeConn(conn)
	}

}
