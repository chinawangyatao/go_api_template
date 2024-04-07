// Package middleware 鉴权中间件
package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go_admin/common"
	"go_admin/server/entity"
	"go_admin/utils"
	"strings"
)

type userInfo struct {
	entity.JwtUser
	jwt.RegisteredClaims
}

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
		var token = p[1]
		jwtUser, err := ParseToken(token)
		if err != nil {
			utils.Failed(c, 401, "token 格式错误！")
			utils.L.Error("ParseToken error!")
			c.Abort()
			return
		}
		c.Set(ContextKeyUserObj, jwtUser)
		c.Next()
	}
}

// ParseToken 解析 jwt
func ParseToken(tokenString string) (*entity.JwtUser, error) {
	if tokenString == "" { // 如果为空返回不存在
		return nil, errors.New(common.ErrAbsent)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return common.Secret, nil
	})

	if token == nil {
		return nil, errors.New(common.ErrInvalid)
	}

	users := userInfo{}
	_, err = jwt.ParseWithClaims(tokenString, &users, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
		}
		return common.Secret, nil
	})

	if err != nil {
		return nil, err
	}
	return &users.JwtUser, err
}
