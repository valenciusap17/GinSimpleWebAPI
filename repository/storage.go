package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/valenciusap17/GoT_Auth/models"
)

type Storage interface {
	CreateUser(*models.User) (*models.User, error)
	ReadUserById(uuid.UUID) (*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
	DeleteUser(uuid.UUID) (*models.User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=primayudha sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) CreateUser(*models.User) (*models.User, error) {
	return nil, nil
}
func (s *PostgresStore) ReadUserById(uuid.UUID) (*models.User, error) {
	return nil, nil

}
func (s *PostgresStore) UpdateUser(*models.User) (*models.User, error) {
	return nil, nil
}
func (s *PostgresStore) DeleteUser(uuid.UUID) (*models.User, error) {
	return nil, nil
}
