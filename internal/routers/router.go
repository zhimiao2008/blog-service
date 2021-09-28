package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-grogramming-tour-book/blog-service/docs"
	"github.com/go-grogramming-tour-book/blog-service/global"
	"github.com/go-grogramming-tour-book/blog-service/internal/middleware"
	"github.com/go-grogramming-tour-book/blog-service/internal/routers/api"
	"github.com/go-grogramming-tour-book/blog-service/internal/routers/api/v1"
	"github.com/go-grogramming-tour-book/blog-service/pkg/limiter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     1,
		Quantum:      1,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.MetricsMonitor())
	r.Use(middleware.Translations())
	r.Use(middleware.AppInfo())
	r.Use(middleware.Tracing())

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/panic", api.GeneratePanic)
	r.GET("/timeout", api.GenerateTimeout)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// expose prometheus metrics接口

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.Static("/static", string(http.Dir(global.AppSetting.UploadSavePath)))

	r.GET("/auth", api.GetAuth)
	r.GET("/info", api.GetInfo)

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		// 创建标签
		apiv1.POST("/tags", tag.Create)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiv1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiv1.GET("/tags", tag.List)

		// 创建文章
		apiv1.POST("/articles", article.Create)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiv1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiv1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiv1.GET("/articles", article.List)
	}
	return r
}
