package router

import (
	"simple_bank/internal/router/admin"
	"simple_bank/internal/router/client"
	"simple_bank/internal/router/public"
)

type RouterGroup struct {
	Client client.ClientRouterGroup
	Admin  admin.AdminRouterGroup
	Public public.PublicRouterGroup
}

var RouterGroupApp = new(RouterGroup)
