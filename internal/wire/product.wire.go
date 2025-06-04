//go:build wireinject

package wire

import (
	"bds-square-backend/internal/controller"
	"bds-square-backend/internal/repo"
	"bds-square-backend/internal/service"

	"github.com/google/wire"
)

func InitProductRouterHanlder() (*controller.ProductController, error) {
	wire.Build(
		repo.NewProductRepository,
		service.NewProductService,
		controller.NewProductController,
	)
	return new(controller.ProductController), nil
}
