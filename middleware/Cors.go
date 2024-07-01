package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// 允许跨域的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		// 允许跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		// 允许跨域请求的响应头
		c.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers, Content-Type")
		// 允许跨域请求携带cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		method := c.Request.Method
		// OPTIONS方法用于获取目的资源所支持的通信选项
		if method == "OPTIONS" {
			// 如果是OPTIONS请求，直接返回空响应
			c.AbortWithStatus(http.StatusNoContent)
		}

	}
}
