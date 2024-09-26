package routes

import (
	"gateway-service/internal/router/manager"
	"gateway-service/internal/router/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
