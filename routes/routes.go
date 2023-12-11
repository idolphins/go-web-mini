package routes

import (
	"fmt"
	"osstp-go-hive/app/admin/middleware"
	"osstp-go-hive/config"
	_ "osstp-go-hive/docs"
	"osstp-go-hive/global"
	mobile_middleware "osstp-go-hive/mobile/admin/middleware"
	pkg_middleware "osstp-go-hive/pkg/middleware"
	"osstp-go-hive/routes/app/admin"
	mobile_admin "osstp-go-hive/routes/mobile/admin"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化
func InitRoutes() *gin.Engine {
	//设置模式
	gin.SetMode(config.Config.System.Mode)
	r := gin.Default()

	// 启用限流中间件
	// 默认每50毫秒填充一个令牌，最多填充200个
	fillInterval := time.Duration(config.Config.RateLimit.FillInterval)
	capacity := config.Config.RateLimit.Capacity
	r.Use(pkg_middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))

	// 启用全局跨域中间件
	r.Use(pkg_middleware.CORSMiddleware())

	// 启用操作日志中间件
	r.Use(pkg_middleware.OperationLogMiddleware())

	// http://localhost:8088/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 初始化JWT认证中间件
	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		global.ZLog.Panicf("WEB初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("WEB初始化JWT中间件失败：%v", err))
	}

	// 初始化Web端路由组
	appGroup := r.Group("/" + config.Config.System.WebUrlPathPrefix)
	admin.InitAdminRoutes(appGroup, authMiddleware)

	mobileMiddleware, err := mobile_middleware.InitMobileAuth()
	if err != nil {
		global.ZLog.Panicf("Mobile初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("Mobile初始化JWT中间件失败：%v", err))
	}
	// 初始化Mobile端路由组
	mobileGroup := r.Group("/" + config.Config.System.MobileUrlPathPrefix)
	mobile_admin.InitMobileRoutes(mobileGroup, mobileMiddleware)

	global.ZLog.Info("初始化路由完成！")
	return r
}
