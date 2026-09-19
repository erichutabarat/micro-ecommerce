package data

import (
	"time"
	"context"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	Conn *sql.DB
}

func NewPostgresRepository(db *sql.DB) *Repository {
	return &Repository{
		Conn: db,
	}
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
	FirstName    sql.NullString
	LastName     sql.NullString
}

func (repo *Repository) Authenticate(email, password string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, email, password_hash, first_name, last_name FROM users WHERE email = $1`
	
	var user User
	err := repo.Conn.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("user not found")
		}
		log.Printf("Database error during authentication: %v", err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}

	return true, nil
}