package mobile_admin

import (
	"osstp-go-hive/mobile/admin/route"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitMobileRoutes(apiGroup *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	route.MobileInitBaseRoutes(apiGroup, authMiddleware)
	route.MobileInitUserRoutes(apiGroup, authMiddleware)

}
