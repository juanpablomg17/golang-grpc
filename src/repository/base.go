package repository

import (
	"context"
	"protobuf-grpc/src/model"
)

type RepositoryBase[T model.ModelBase] interface {
	GetItem(ctx context.Context, id string) (*T, error)
	SetItem(ctx context.Context, item *T) error
}
