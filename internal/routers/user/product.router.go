package user

import (
	"bds-square-backend/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productController, _ := wire.InitProductRouterHanlder()
	// public router
	productRouterPublic := Router.Group("/user/products")
	{
		productRouterPublic.GET("", productController.GetList)
		productRouterPublic.GET("/:id", productController.GetByID)
		productRouterPublic.POST("", productController.Create)
		productRouterPublic.PUT("/:id", productController.Update)
		productRouterPublic.DELETE("/:id", productController.Delete)
	}

	// private router
	productRouterPrivate := Router.Group("/user/products")
	// productRouterPrivate.Use(limiter())
	// productRouterPrivate.Use(authen())
	// productRouterPrivate.Use(permission())
	{
		productRouterPrivate.GET("/active_product")
	}
}
