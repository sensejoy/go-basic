package controller

import (
	"encoding/json"
	"go-basic/lib"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func InitRoute(server *gin.Engine) {
	server.Use(requestInit(), gin.Recovery(), log())
	server.GET("/", sample)
	server.GET("/user/get", getUser)
	server.POST("/user/add", addUser)
}

func requestInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("start", time.Now().UnixNano())
		c.Next()
	}
}

func log() gin.HandlerFunc {
	return func(c *gin.Context) {
		end := time.Now().UnixNano()
		start, _ := c.Get("start")
		cost := (int)((end - start.(int64)) / 1000000)
		r := c.Request
		header, _ := json.Marshal(r.Header)
		cookie, _ := json.Marshal(r.Cookies())
		lib.Logger.Info("",
			zap.String("request_id", c.GetHeader("X-REQUEST-ID")),
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.String("param", r.PostForm.Encode()),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("header", string(header)),
			zap.String("cookie", string(cookie)),
			zap.Int("cost", cost),
		)
	}
}

func sample(c *gin.Context) {
	genResult(c, 0, "ok", nil)
}

func genResult(c *gin.Context, errno int, errmsg string, data interface{}) {
	c.JSON(200, gin.H{
		"errno":  errno,
		"errmsg": errmsg,
		"data":   data,
	})
}
