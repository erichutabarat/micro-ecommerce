package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type LoginRequest struct {
	Email string
	Password string
}

type LoginResponse struct {
	Success bool
	Message string
	Token string
}

type AuthService struct {

}

func (t *AuthService) Login(req *LoginRequest, res *LoginResponse) error {
	log.Printf("Receive request for: %s", req.Email)

	if req.Email == "test@example.com" && req.Password == "123" {
		res.Success = true
		res.Message = "Authentication successful"
		res.Token = "token-123"
		return nil
	}

	res.Success = false
	res.Message = "Authentication failed"
	return errors.New("Authentication failed")
}

func main() {
	auth := new(AuthService)

	err := rpc.Register(auth)
	if err != nil {
		log.Fatal("Error registering rpc service: %s", err)
	}

	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Error listening on port 5001: %v", err)
	}
	defer listener.Close()

	log.Println("Auth RPC Service is running on port 5001...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}