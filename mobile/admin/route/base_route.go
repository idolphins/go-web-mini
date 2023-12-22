package route

import (
	"osstp-go-hive/app/admin/controller"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册基础路由
func MobileInitBaseRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	router := r.Group("/base")
	{
		// 登录登出刷新token无需鉴权
		router.POST("/login", authMiddleware.LoginHandler)
		router.POST("/logout", authMiddleware.LogoutHandler)
		router.POST("/refreshToken", authMiddleware.RefreshHandler)
	}

	baseController := controller.NewBaseController()
	/// 用户注册
	{
		//无需鉴权和JWT认证
		router.POST("/register", baseController.Register) // 用户注册账号
		// userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		// userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		// userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		// userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		// userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		// userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		// userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	return r
}
