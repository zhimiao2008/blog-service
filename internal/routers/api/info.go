package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-grogramming-tour-book/blog-service/pkg/app"
)

func GetInfo(c *gin.Context) {
	response := app.NewResponse(c)

	appName := c.GetString("app_name")
	appVersion := c.GetString("app_version")

	response.ToResponse(gin.H{
		"appName":    appName,
		"appVersion": appVersion,
	})
	return

}
