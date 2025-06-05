package routers

import (
	"bds-square-backend/internal/routers/admin"
	"bds-square-backend/internal/routers/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
