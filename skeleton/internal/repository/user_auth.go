package repository

import (
	"context"
	"github.com/kercylan98/minotaur/skeleton/internal/repository/mapped"
)

type UserAuthRepository interface {
	Repository[*mapped.UserAuthMapped]

	CreateUserAuth(ctx context.Context, dto *CreateUserAuthDTO) (err error)

	DeleteUserAuthByUserId(ctx context.Context, dto *DeleteUserAuthByUserIdDTO) (err error)
}

type CreateUserAuthDTO struct {
	UserId   string              `json:"user_id"`   // 用户 ID
	AuthType mapped.UserAuthType `json:"auth_type"` // 授权类型
	Account  string              `json:"account"`   // 授权账号，通常为邮箱、手机号、第三方账号等
	AuthData string              `json:"auth_data"` // 授权数据，如加密后的密码、第三方授权等
}

type DeleteUserAuthByUserIdDTO struct {
	UserId string `json:"user_id"` // 用户 ID
}
