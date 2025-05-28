package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(authen())
	// adminRouterPrivate.Use(permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
