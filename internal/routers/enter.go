package routers

import (
	"bds-square-backend/internal/routers/manage"
	"bds-square-backend/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
