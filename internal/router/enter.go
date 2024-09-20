package routes

import (
	"github.com/bhtoan2204/go-practice.git/internal/router/manager"
	"github.com/bhtoan2204/go-practice.git/internal/router/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
