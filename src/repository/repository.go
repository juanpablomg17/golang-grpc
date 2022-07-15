package repository

import (
	"context"
	"protobuf-grpc/src/model"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*model.Student, error)
	SetStudent(ctx context.Context, student *model.Student) error
	GetExam(ctx context.Context, id string) (*model.Exam, error)
	SetExam(ctx context.Context, exam *model.Exam) error
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func SetStudent(ctx context.Context, student *model.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*model.Student, error) {
	return implementation.GetStudent(ctx, id)
}
