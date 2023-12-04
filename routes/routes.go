package routes

import (
	"fmt"
	"osstp-go-hive/app/admin/middleware"
	"osstp-go-hive/config"
	"osstp-go-hive/global"
	"osstp-go-hive/routes/admin"
	"osstp-go-hive/routes/shop"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化
func InitRoutes() *gin.Engine {
	//设置模式
	gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()
	// 创建不带中间件的路由:
	// r := gin.New()
	// r.Use(gin.Recovery())

	// 启用限流中间件
	// 默认每50毫秒填充一个令牌，最多填充200个
	fillInterval := time.Duration(config.Conf.RateLimit.FillInterval)
	capacity := config.Conf.RateLimit.Capacity
	r.Use(middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))

	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 启用操作日志中间件
	r.Use(middleware.OperationLogMiddleware())

	// 初始化JWT认证中间件
	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		global.ZLog.Panicf("初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}

	// 路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 初始化Admin路由组
	admin.InitAdminRoutes(apiGroup, authMiddleware)
	// 初始化shop路由组
	shop.InitShopRoutes(apiGroup, authMiddleware)

	global.ZLog.Info("初始化路由完成！")
	return r
}
