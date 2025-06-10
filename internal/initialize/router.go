package initialize

import (
	"bds-square-backend/global"
	"bds-square-backend/internal/middlewares"
	"bds-square-backend/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global
	r.Use(middlewares.CORSMiddleware())

	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/checkStatus") // tracking monitor

		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)

		adminRouter.InitUserRouter(MainGroup)
		adminRouter.InitAdminRouter(MainGroup)
	}

	return r
}
