package main

import (
	"database/sql"
	"errors"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"

	"auth/data" 
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Success bool
	Message string
	Token   string
}

type AuthService struct {
	Repo *data.Repository
}

func (t *AuthService) Login(req *LoginRequest, res *LoginResponse) error {
	log.Printf("Receive request for: %s", req.Email)

	valid, err := t.Repo.Authenticate(req.Email, req.Password)
	if err != nil || !valid {
		res.Success = false
		res.Message = "Authentication failed"
		return errors.New("authentication failed")
	}

	res.Success = true
	res.Message = "Authentication successful"
	res.Token = "token-123" 
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, relying on system environment variables")
	}

	log.Println("Starting auth service...")

	db := connectToDB()
	if db == nil {
		log.Panic("Can not connect to Postgres!")
	}
	defer db.Close()

	repo := data.NewPostgresRepository(db)
	auth := &AuthService{
		Repo: repo,
	}

	err = rpc.Register(auth)
	if err != nil {
		log.Fatalf("Error registering rpc service: %v", err)
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Panic("Cant read dsn")
	}
	log.Printf("dsn= %s", dsn)
	counts := 0
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready... sleeping 2 seconds")
			time.Sleep(2 * time.Second)
			counts++
		} else {
			log.Println("Connected to Postgres database!")
			return connection
		}

		if counts > 10 {
			log.Println("Max connection attempts reached. Exiting.")
			return nil
		}
	}
}