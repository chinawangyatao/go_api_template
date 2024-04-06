// Package middleware 鉴权中间件
package middleware

import (
	"github.com/gin-gonic/gin"
	"go_admin/utils"
	"strings"
)

var ContextKeyUserObj = "authUserObj"

func AuthJWT() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization") // 获取请求头里面的 Authorization
		if authorization == "" {
			utils.Failed(c, int(utils.ApiCode.NOAUTH), utils.ApiCode.GetMessage(utils.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		p := strings.SplitN(authorization, " ", 2)
		if !(len(p) == 2 && p[0] == "Bearer") {
			utils.Failed(c, int(utils.ApiCode.AUTHFORMATERROR), utils.ApiCode.GetMessage(utils.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}

		//	todo 校验 token
		var token = "token"
		c.Set(ContextKeyUserObj, token)
		c.Next()
	}
}
