package admin

import (
	"osstp-go-hive/app/admin/route"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(apiGroup *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	// 用户登录、退出，新用户注册
	route.InitBaseRoutes(apiGroup, authMiddleware)         // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	route.InitUserRoutes(apiGroup, authMiddleware)         // 注册用户路由, jwt认证中间件,casbin鉴权中间件
	route.InitRoleRoutes(apiGroup, authMiddleware)         // 注册角色路由, jwt认证中间件,casbin鉴权中间件
	route.InitMenuRoutes(apiGroup, authMiddleware)         // 注册菜单路由, jwt认证中间件,casbin鉴权中间件
	route.InitApiRoutes(apiGroup, authMiddleware)          // 注册接口路由, jwt认证中间件,casbin鉴权中间件
	route.InitOperationLogRoutes(apiGroup, authMiddleware) // 注册操作日志路由, jwt认证中间件,casbin鉴权中间件

}
