package system

import (
	v1 "github.com/coder-th/ds-tools/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)            // 登录
		baseRouter.POST("captcha", baseApi.Captcha)        // 获取验证码
		baseRouter.POST("adminRegister", baseApi.Register) // 管理员注册账号
	}
	return baseRouter
}
