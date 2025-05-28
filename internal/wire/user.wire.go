//go:build wireinject

package wire

import (
	"bds-square-backend/internal/controller"
	"bds-square-backend/internal/repo"
	"bds-square-backend/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHanlder() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}