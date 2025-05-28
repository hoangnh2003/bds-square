package user

import (
	"bds-square-backend/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHanlder()
	// public router
	userRouterPublic := Router.Group("/user")
	{
		// userRouterPublic.GET("/getUser", controller.NewUserController().GetUserByID)
		userRouterPublic.GET("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(authen())
	// userRouterPrivate.Use(permission())

	userRouterPrivate.GET("/get_info")
}
