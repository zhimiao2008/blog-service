package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-grogramming-tour-book/blog-service/internal/routers/api/v1"
)


func  GeneratePanic(c *gin.Context) {
	r := &v1.Article{}
	r = nil
	r.Create(c)
}