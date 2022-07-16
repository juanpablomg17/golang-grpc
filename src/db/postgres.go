package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}
