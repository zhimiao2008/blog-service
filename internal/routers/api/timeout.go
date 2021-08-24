package api

import (
	"github.com/gin-gonic/gin"
	"time"
)

func GenerateTimeout(c *gin.Context) {
	time.Sleep(60 * time.Second)
}
