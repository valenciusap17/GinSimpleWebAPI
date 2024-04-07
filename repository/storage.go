package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/valenciusap17/GoT_Auth/models"
)

type Storage interface {
	CreateUser(*models.User) (*models.User, error)
	ReadAllUser() ([]*models.User, error)
	ReadUserById(*uuid.UUID) (*models.User, error)
	UpdateUser(*models.CreateUserRequest, *uuid.UUID) (*models.User, error)
	DeleteUser(*uuid.UUID) (*models.User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func getEnvVariables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func NewPostgresStore() (*PostgresStore, error) {
	user := getEnvVariables("POSTGRES_USER")
	port, err := strconv.Atoi(getEnvVariables("POSTGRES_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	DB := getEnvVariables("POSTGRES_DB")
	password := getEnvVariables("POSTGRES_PASSWORD")
	host := getEnvVariables("POSTGRES_HOST")

	connStr := fmt.Sprintf("port=%d user=%s dbname=%s password=%s host=%s sslmode=disable",
		port, user, DB, password, host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateUserTable()
}

func (s *PostgresStore) CreateUserTable() error {
	query := `create table if not exists users (
		id uuid primary key default gen_random_uuid(),
		username varchar(55),
		password varchar(255),
		firstname varchar(55),
		lastname varchar(55),
		createdat timestamp,
		updatedat timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateUser(data *models.User) (*models.User, error) {
	queryResponse := new(models.User)

	query := `insert into users 
		(username, password, firstname, lastname, createdat)
		values ($1, $2, $3, $4, $5)
		returning id, username, password, firstname, lastname, createdat
	`
	err := s.db.QueryRow(
		query,
		data.Username,
		data.Password,
		data.FirstName,
		data.LastName,
		data.CreatedAt).Scan(&queryResponse.ID,
		&queryResponse.Username,
		&queryResponse.Password,
		&queryResponse.FirstName,
		&queryResponse.LastName,
		&queryResponse.CreatedAt)
	if err != nil {
		return nil, err
	}

	return queryResponse, nil
}

func (s *PostgresStore) ReadAllUser() ([]*models.User, error) {
	query := `select * from users`

	response, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	queryResponse := []*models.User{}

	for response.Next() {
		eachUser := new(models.User)
		err := response.Scan(
			&eachUser.ID,
			&eachUser.Username,
			&eachUser.Password,
			&eachUser.FirstName,
			&eachUser.LastName,
			&eachUser.CreatedAt,
			&eachUser.UpdatedAt)

		if err != nil {
			return nil, err
		}

		queryResponse = append(queryResponse, eachUser)
	}
	return queryResponse, nil
}

func (s *PostgresStore) ReadUserById(id *uuid.UUID) (*models.User, error) {
	response := &models.User{}

	query := `select * from users where id = $1`
	err := s.db.QueryRow(query, id).Scan(
		&response.ID,
		&response.Username,
		&response.Password,
		&response.FirstName,
		&response.LastName,
		&response.CreatedAt)

	if err != nil {
		return nil, err
	}

	return response, nil

}
func (s *PostgresStore) UpdateUser(data *models.CreateUserRequest, id *uuid.UUID) (*models.User, error) {
	response := &models.User{}

	query := `udpate users 
		set username = $1, password $2, firstname = $3, lastname = $4, updatedat = $5
		where id = $6
		returning id, username, password, firstname, lastname, createdat, updatedat
	`

	err := s.db.QueryRow(query,
		data.Username,
		data.Password,
		data.FirstName,
		data.LastName,
		time.Now().UTC(),
		id).Scan(
		&response.ID,
		&response.Username,
		&response.Password,
		&response.FirstName,
		&response.LastName,
		&response.CreatedAt,
		&response.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return response, nil
}
func (s *PostgresStore) DeleteUser(id *uuid.UUID) (*models.User, error) {
	response := &models.User{}

	query := `delete from users 
		where id = $1
		returning id, username, password, firstname, lastname, createdat, updatedat
	`

	err := s.db.QueryRow(query, id).Scan(
		&response.ID,
		&response.Username,
		&response.Password,
		&response.FirstName,
		&response.LastName,
		&response.CreatedAt,
		&response.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return response, nil
}
