package pkg_tool

import (
	"osstp-go-hive/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func TrimPrefixPath(c *gin.Context) string {
	var obj string
	if strings.Contains(c.FullPath(), "/"+config.Config.System.WebUrlPathPrefix) {
		obj = strings.TrimPrefix(c.FullPath(), "/"+config.Config.System.WebUrlPathPrefix)
	} else if strings.Contains(c.FullPath(), "/"+config.Config.System.MobileUrlPathPrefix) {
		obj = strings.TrimPrefix(c.FullPath(), "/"+config.Config.System.MobileUrlPathPrefix)
	} else {
		obj = c.FullPath()
	}
	return obj
}
