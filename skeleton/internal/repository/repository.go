package repository

import "context"

type Repository[Model any] interface {
	// InitTable 初始化表结构
	InitTable(ctx context.Context, model Model) (err error)
}
