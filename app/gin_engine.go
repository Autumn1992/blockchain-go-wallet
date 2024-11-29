package app

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"runtime"
	"time"
	"walletserver/log"
	"walletserver/models/response"
)

var skippedPathList = map[string]struct{}{
	"/adminapi/bill/heartbeat": {},
}

func SkippedPath(path string) bool {
	if _, ok := skippedPathList[path]; ok {
		return true
	}
	return false
}
func NewGinEngine() *gin.Engine {
	gin.DefaultErrorWriter = log.ErrorWriter
	gin.DefaultWriter = log.InfoWriter
	//去掉默认的打印
	engine := gin.New()
	//engine := gin.Default()
	//engine.Use(gin.Recovery())
	engine.Use(gin.CustomRecovery(handleRecovery()))
	engine.Use(logger())
	//engine.Use(middle.AddTraceId())

	return engine
}

// HandleRecovery API错误捕捉
func handleRecovery() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		// 捕捉错误，进行返回码处理
		response.Fail(c, fmt.Sprintf("服务器recovery异常 %v", err))
		//获取异常报错堆栈
		buf := make([]byte, 1<<20)
		runtime.Stack(buf, false)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		bodyBytes, _ := c.GetRawData()

		c.Set("reqstarttime", time.Now().UnixMilli())
		// Process request

		// Log only when path is not being skipped
		param := gin.LogFormatterParams{
			Request: c.Request,
		}
		//log.Error(c.Keys)
		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		var postparam = string(bodyBytes)
		var printLog = !SkippedPath(param.Path)
		c.Set("printLog", printLog)
		if printLog {
			log.Info(fmt.Sprintf("Request %s %3d %13v %s %s[%s]",
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				param.Path), c.Request.URL.Query(), postparam)
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	}
}
