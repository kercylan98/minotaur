package services

import (
	"context"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/google/uuid"
	"github.com/kercylan98/minotaur/skeleton/internal/repository"
	"github.com/kercylan98/minotaur/skeleton/internal/repository/mapped"
	"github.com/kercylan98/minotaur/skeleton/internal/services/bo"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
)

func NewUserService(dtmServer string, user repository.UserRepository, userAuth repository.UserAuthRepository) *UserService {
	return &UserService{
		dtmServer: dtmServer,
		user:      user,
		userAuth:  userAuth,
	}
}

type UserService struct {
	fiberApp  *fiber.App
	dtmServer string

	user     repository.UserRepository
	userAuth repository.UserAuthRepository
}

func (s *UserService) Init() {
	s.fiberApp.Post("/api/v1/user/create", s.onCreateUser)
	s.fiberApp.Post("/api/v1/user/create/compensate", s.onCreateUserCompensate)
	s.fiberApp.Post("/api/v1/user/auth/create", s.onCreateUserAuth)
	s.fiberApp.Post("/api/v1/user/auth/create/compensate", s.onCreateUserAuthCompensate)
}

func (s *UserService) CreateUserWithAccount(ctx context.Context, request bo.CreateUserWithAccountBO) (err error) {
	saga := dtmcli.NewSaga(s.dtmServer, uuid.NewString()).
		Add("/v1/user/create/action", "/v1/user/create/compensate", request).
		Add("/v1/user/auth/create/action", "/v1/user/auth/create/compensate", request)
	return saga.Submit()
}

func (s *UserService) onCreateUser(ctx *fiber.Context) error {
	//barrier := infoFromContext(ctx)
	var req bo.CreateUserWithAccountBO
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	return s.user.CreateUser(context.Background(), &repository.CreateUserDTO{
		UserId: req.UserId,
	})

}

func (s *UserService) onCreateUserCompensate(ctx *fiber.Context) error {
	var req bo.CreateUserWithAccountBO
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	return s.user.DeleteUserById(context.Background(), &repository.DeleteUserByIdDTO{
		UserId: req.UserId,
	})
}

func (s *UserService) onCreateUserAuth(ctx *fiber.Context) error {
	var req bo.CreateUserWithAccountBO
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	return s.userAuth.CreateUserAuth(context.Background(), &repository.CreateUserAuthDTO{
		UserId:   req.UserId,
		AuthType: mapped.AccountUserAuthType,
		Account:  req.Account,
		AuthData: req.Password,
	})
}

func (s *UserService) onCreateUserAuthCompensate(ctx *fiber.Context) error {
	var req bo.CreateUserWithAccountBO
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	return s.userAuth.DeleteUserAuthByUserId(context.Background(), &repository.DeleteUserAuthByUserIdDTO{
		UserId: req.UserId,
	})
}

func infoFromContext(c *fiber.Context) *dtmcli.BranchBarrier {
	info := dtmcli.BranchBarrier{
		TransType: c.Query("trans_type"),
		Gid:       c.Query("gid"),
		BranchID:  c.Query("branch_id"),
		Op:        c.Query("op"),
	}
	return &info
}
