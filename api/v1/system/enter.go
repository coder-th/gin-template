package system

import "github.com/coder-th/ds-tools/server/service"

type ApiGroup struct {
	JwtApi
	BaseApi
}

var (
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
