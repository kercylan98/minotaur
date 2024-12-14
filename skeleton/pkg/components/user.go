package components

import modelsv1 "github.com/kercylan98/minotaur/skeleton/common/models/v1"

type UserComponent interface {
	CreateUser(userinfo *modelsv1.CreateUserRequest) (*modelsv1.CreateUserResponse, error)
}
