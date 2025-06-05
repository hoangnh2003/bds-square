package admin

import (
	"bds-square-backend/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHanlder()

	// public router
	userRouterPublic := Router.Group("/admin/user")
	{
		userRouterPublic.POST("/register", userController.Register)
	}

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(authen())
	// userRouterPrivate.Use(permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
