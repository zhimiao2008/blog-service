package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-grogramming-tour-book/blog-service/global"
)

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", global.AppSetting.AppName)
		c.Set("app_version", global.AppSetting.AppVersion)
		c.Next()
	}
}
