package log

import (
	"en-lang-bk/global"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)

		global.LOG.Info(
			path,
			"  "+fmt.Sprintf("%d", c.Writer.Status()),
			"  "+c.Request.Method+"  ",
			//zap.String("path", path),
			query+"  ",
			c.ClientIP()+"  ",
			c.Request.UserAgent()+"  ",
			c.Errors.ByType(gin.ErrorTypePrivate).String()+"  ",
			cost,
		)
	}
}
