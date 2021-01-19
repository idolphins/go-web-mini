package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-lim/controller"
)

func InitMenuRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	menuController := controller.NewMenuController()
	router := r.Group("/menu")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.GET("/tree", menuController.GetMenuTree)
		router.GET("/all/:roleId", menuController.GetAllMenuByRoleId)
		router.GET("/list", menuController.GetMenus)
		router.POST("/create", menuController.CreateMenu)
		router.PATCH("/update/:menuId", menuController.UpdateMenuById)
		router.DELETE("/delete/batch", menuController.BatchDeleteMenuByIds)
	}

	return r
}
