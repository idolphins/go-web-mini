package route

import (
	"osstp-go-hive/app/admin/controller"
	pkg_middleware "osstp-go-hive/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册用户路由
func MobileInitUserRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	userController := controller.NewUserController()
	router := r.Group("/user")

	router.Use(authMiddleware.MiddlewareFunc())

	router.Use(pkg_middleware.CasbinMiddleware())
	{
		router.POST("/info", userController.GetUserInfo)
		router.GET("/list", userController.GetUsers)
		router.PUT("/changePwd", userController.ChangePwd)
		router.POST("/create", userController.CreateUser)
		router.PATCH("/update/:userId", userController.UpdateUserById)
		router.DELETE("/delete/batch", userController.BatchDeleteUserByIds)
	}
	return r
}
