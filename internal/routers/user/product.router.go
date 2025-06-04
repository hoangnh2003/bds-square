package user

import (
	"bds-square-backend/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productController, _ := wire.InitProductRouterHanlder()
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/list-product", productController.GetList)
		productRouterPublic.GET("/create-product", productController.Create)
		productRouterPublic.GET("/update-product", productController.Update)
		productRouterPublic.GET("/get-product", productController.GetByID)
		productRouterPublic.GET("/delete-product", productController.Delete)
	}

	// private router
}
