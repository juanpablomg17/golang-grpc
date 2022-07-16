package repository

import (
	"context"
	"protobuf-grpc/src/model"
)

type StudentRepository interface {
	RepositoryBase[model.Student]
}

var studentImplemtation StudentRepository

func SetStudentRepository(repo StudentRepository) {
	studentImplemtation = repo
}

func GetStudent(ctx context.Context, id string) (*model.Student, error) {
	return studentImplemtation.GetItem(ctx, id)
}

func SetStudent(ctx context.Context, student *model.Student) error {
	return studentImplemtation.SetItem(ctx, student)
}
