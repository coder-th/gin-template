package system

import (
	v1 "github.com/coder-th/ds-tools/server/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("changePassword", baseApi.ChangePassword) // 用户修改密码
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)       // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)        // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)        // 设置自身信息
		userRouter.POST("resetPassword", baseApi.ResetPassword)   // 充值密码
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
	}
}
