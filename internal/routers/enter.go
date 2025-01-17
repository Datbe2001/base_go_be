package routers

import (
	"base_go_be/internal/routers/manager"
	"base_go_be/internal/routers/user"
)

type RouterGroup struct {
	User    user.UsersRouterGroup
	Manager manager.ManagersRouterGroup
}

var RouterGroupApp = new(RouterGroup)
