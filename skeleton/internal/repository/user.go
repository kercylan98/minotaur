package repository

import (
	"context"
	"github.com/kercylan98/minotaur/skeleton/internal/repository/mapped"
)

type UserRepository interface {
	Repository[*mapped.UserMapped]

	// CreateUser 创建用户
	CreateUser(ctx context.Context, dto *CreateUserDTO) (err error)

	// DeleteUserById 删除用户
	DeleteUserById(ctx context.Context, dto *DeleteUserByIdDTO) (err error)
}

type CreateUserDTO struct {
	UserId string
}

type DeleteUserByIdDTO struct {
	UserId string
}
