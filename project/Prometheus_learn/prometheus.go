package main

import (
	"math/rand"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 自定义业务状态码 Counter 指标
var statusCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_response_status_count",
	},
	[]string{"method", "path", "status"},
)

func initRegistry() *prometheus.Registry {
	// 创建一个 registry
	reg := prometheus.NewRegistry()

	// 添加 Go 编译信息
	reg.MustRegister(collectors.NewBuildInfoCollector())
	// Go runtime metrics
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(
			collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")},
		),
	))

	// 注册自定义的业务指标
	reg.MustRegister(statusCounter)

	return reg
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		// mock 业务逻辑，异常情况下返回 status = 1
		status := 0
		if rand.Intn(10)%3 == 0 {
			status = 1
		}

		// 记录
		statusCounter.WithLabelValues(
			c.Request.Method,
			c.Request.URL.Path,
			strconv.Itoa(status),
		).Inc()

		c.JSON(200, gin.H{
			"status":  status,
			"message": "pong",
		})
	})

	reg := initRegistry()

	// 对外提供 /metrics 接口，支持 prometheus 采集
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{Registry: reg},
	)))

	r.Run(":8083")
}
