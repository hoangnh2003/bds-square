// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"bds-square-backend/internal/controller"
	"bds-square-backend/internal/repo"
	"bds-square-backend/internal/service"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from product.wire.go:

func InitProductRouterHanlder() (*controller.ProductController, error) {
	iProductRepository := repo.NewProductRepository()
	iProductService := service.NewProductService(iProductRepository)
	productController := controller.NewProductController(iProductService)
	return productController, nil
}

// Injectors from user.wire.go:

func InitUserRouterHanlder() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
