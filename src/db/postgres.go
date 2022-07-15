package db

import (
	"context"
	"database/sql"
	"log"
	"protobuf-grpc/src/model"

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

func (repo PostgresRepository) SetStudent(ctx context.Context, student *model.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)

	if err != nil {
		log.Println("In SetStudent implementation repo", err)
		return err
	}

	return nil
}

func (repo PostgresRepository) GetStudent(ctx context.Context, id string) (*model.Student, error) {

	var err error

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM students WHERE id = $1", id)

	if err != nil {
		log.Println("In GetStudent implementation repo", err)
		return nil, err
	}

	var student = &model.Student{}

	if err != nil {
		log.Println("In GetStudent implementation repo", err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
	}

	return student, nil

}
