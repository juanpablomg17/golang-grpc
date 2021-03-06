package repository

import (
	"context"
	"protobuf-grpc/src/model"
)

type ExamRepository interface {
	RepositoryBase[model.Exam]
	SetQuestion(ctx context.Context, question *model.Question) error
	GetStudentsPerExams(ctx context.Context, examId string) ([]*model.Student, error)
	SetEnrollment(ctx context.Context, enrollment *model.Enrollment) error
}

var implementationExam ExamRepository

func SetExamRepository(repo ExamRepository) {
	implementationExam = repo
}

func SetExam(ctx context.Context, exam *model.Exam) error {
	return implementationExam.SetItem(ctx, exam)
}

func GetExam(ctx context.Context, id string) (*model.Exam, error) {
	return implementationExam.GetItem(ctx, id)
}
