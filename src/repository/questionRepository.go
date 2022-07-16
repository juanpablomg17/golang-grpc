package repository

import (
	"context"
	"protobuf-grpc/src/model"
)

type QuestionRepository interface {
	RepositoryBase[model.Question]
}

var implementationQuestion QuestionRepository

func SetQuestionRepository(repo QuestionRepository) {
	implementationQuestion = repo
}

func SetQuestion(ctx context.Context, question *model.Question) error {
	return implementationQuestion.SetItem(ctx, question)
}

func GetQuestion(ctx context.Context, id string) (*model.Question, error) {
	return nil, nil
}
