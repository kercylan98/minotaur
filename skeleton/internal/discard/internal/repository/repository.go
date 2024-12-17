package repository

import "context"

type Repository[Model any] interface {
	OnInitialize(ctx context.Context, model Model) (err error)
}
