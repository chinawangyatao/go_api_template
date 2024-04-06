// Package middleware 跨域中间件
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Header", "Content-type,AccessToken,X-CSRF-Token,Authorization,Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		//	放行所有的 options
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		//	处理请求
		context.Next()
	}
}
