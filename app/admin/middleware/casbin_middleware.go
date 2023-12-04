package middleware

import (
	"osstp-go-hive/app/admin/dao"
	"osstp-go-hive/config"
	"osstp-go-hive/global"
	pkg_response "osstp-go-hive/pkg/response"

	"github.com/gin-gonic/gin"

	"strings"
	"sync"
)

var checkLock sync.Mutex

// Casbin中间件, 基于RBAC的权限访问控制模型
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := pkg_response.Ctx{Context: c}
		ur := dao.NewUserDao()
		user, err := ur.GetCurrentUser(c)
		if err != nil {
			ctx.Response(401, nil, pkg_response.ResponseMessage{Msg: "用户未登录"})
			c.Abort()
			return
		}
		if user.Status != 1 {
			ctx.Response(401, nil, pkg_response.ResponseMessage{Msg: "当前用户已被禁用"})

			c.Abort()
			return
		}
		// 获得用户的全部角色
		roles := user.Roles
		// 获得用户全部未被禁用的角色的Keyword
		var subs []string
		for _, role := range roles {
			if role.Status == 1 {
				subs = append(subs, role.Keyword)
			}
		}
		// 获得请求路径URL
		//obj := strings.Replace(c.Request.URL.Path, "/"+config.Config.System.UrlPathPrefix, "", 1)
		obj := strings.TrimPrefix(c.FullPath(), "/"+config.Config.System.UrlPathPrefix)
		// 获取请求方式
		act := c.Request.Method

		isPass := check(subs, obj, act)
		if !isPass {
			ctx.Response(401, nil, pkg_response.ResponseMessage{Msg: "没有权限"})

			c.Abort()
			return
		}

		c.Next()
	}
}

func check(subs []string, obj string, act string) bool {
	// 同一时间只允许一个请求执行校验, 否则可能会校验失败
	checkLock.Lock()
	defer checkLock.Unlock()
	isPass := false
	for _, sub := range subs {
		pass, _ := global.CasbinEnforcer.Enforce(sub, obj, act)
		if pass {
			isPass = true
			break
		}
	}
	return isPass
}
