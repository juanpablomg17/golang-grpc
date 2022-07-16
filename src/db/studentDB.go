package db

import (
	"context"
	"log"
	"protobuf-grpc/src/model"
)

type StudentRepository struct {
	PostgresRepository
}

func NewStudentRepository(repo *PostgresRepository) *StudentRepository {
	return &StudentRepository{
		PostgresRepository: *repo,
	}
}

func (repo StudentRepository) SetItem(ctx context.Context, student *model.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)

	if err != nil {
		log.Println("In SetStudent implementation repo", err)
		return err
	}

	return nil
}

func (repo StudentRepository) GetItem(ctx context.Context, id string) (*model.Student, error) {
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
