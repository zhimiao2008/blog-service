package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// 初始化 web_reqeust_total， counter类型指标， 表示接收http请求总次数
var WebRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "web_reqeust_total",
		Help: "Number of hello requests in total",
	},

	// 设置两个标签 请求方法和 路径 对请求总次数在两个
	[]string{"method", "endpoint"},
)

// web_request_duration_seconds，Histogram类型指标，bucket代表duration的分布区间
var WebRequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "web_request_duration_seconds",
		Help:    "web request duration distribution",
		Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 1},
	},
	[]string{"method", "endpoint"},
)

func init() {
	// 注册监控指标
	prometheus.MustRegister(WebRequestTotal)
	prometheus.MustRegister(WebRequestDuration)
}

func MetricsMonitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		// counter类型 metric的记录方式
		WebRequestTotal.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Inc()
		// Histogram类型 metric的记录方式
		WebRequestDuration.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.Request.URL.Path}).Observe(duration.Seconds())
	}
}
