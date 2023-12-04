package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	dao "osstp-go-hive/app/admin/dao"
	"osstp-go-hive/app/admin/middleware"
	"osstp-go-hive/config"
	"osstp-go-hive/global"
	"osstp-go-hive/initialize"
	"osstp-go-hive/routes"
	"syscall"
	"time"
)

func init() {
	// 加载配置文件到全局配置结构体
	config.InitConfig()

	// 初始化日志
	global.ZLog = initialize.InitLogger()

	// 初始化数据库(mysql)
	global.DB = initialize.InitMysql()

	// 初始化casbin策略管理器
	global.CasbinEnforcer = initialize.InitCasbinEnforcer()

	// 初始化Validator数据校验
	global.Validate, global.Trans = initialize.InitValidate()

	// 初始化mysql数据
	initialize.InitData()
}

func main() {

	// 操作日志中间件处理日志时没有将日志发送到rabbitmq或者kafka中, 而是发送到了channel中
	// 这里开启3个goroutine处理channel将日志记录到数据库
	logdao := dao.NewOperationLogDao()
	for i := 0; i < 3; i++ {
		go logdao.SaveOperationLogChannel(middleware.OperationLogChan)
	}

	// 注册所有路由
	r := routes.InitRoutes()

	host := config.Config.System.Host
	port := config.Config.System.Port

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.ZLog.Fatalf("listen: %s\n", err)
		}
	}()

	global.ZLog.Info(fmt.Sprintf("Server is running at %s:%d/%s", host, port, config.Config.System.UrlPathPrefix))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.ZLog.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.ZLog.Fatal("Server forced to shutdown:", err)
	}

	global.ZLog.Info("Server exiting!")

}
