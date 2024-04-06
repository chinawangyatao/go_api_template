// Package jwt jwt 工具类
// 生成 token
// 解析 token
// 获取当前用户信息
package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go_admin/middleware"
	"go_admin/server/entity"
	"time"
)

type userInfo struct {
	entity.JwtUser
	jwt.RegisteredClaims
}

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 24 * 7 * 365

// Secret token 密钥
var Secret = []byte("go-api")
var (
	ErrAbsent  = "token absent 令牌不存在" //令牌不存在
	ErrInvalid = "token invalid 令牌无效" // 令牌无效
)

// GenerateToken 根据用户信息生成 token
func GenerateToken(info entity.SysUsers) (string, error) {
	jwtInfo := entity.JwtUser{
		ID:       info.ID,
		Username: info.Username,
		Nickname: info.Nickname,
		Avatar:   info.Avatar,
		Email:    info.Email,
		Phone:    info.Phone,
		Note:     info.Note,
	}
	c := userInfo{
		JwtUser: jwtInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
			Issuer:    "admin",                                                 //签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signed, err := token.SignedString(Secret)
	return signed, err
}

// ParseToken 解析 jwt
func ParseToken(tokenString string) (*entity.JwtUser, error) {
	if tokenString == "" { // 如果为空返回不存在
		return nil, errors.New(ErrAbsent)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if token == nil {
		return nil, errors.New(ErrInvalid)
	}

	users := userInfo{}
	_, err = jwt.ParseWithClaims(tokenString, &users, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", token.Header["alg"])
		}
		return Secret, nil
	})

	if err != nil {
		return nil, err
	}
	return &users.JwtUser, err
}

// GetUserId 返回用户 Id
func GetUserId(c *gin.Context) (uint, error) {
	u, exist := c.Get(middleware.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("can't get user's id")
	}

	user, ok := u.(*entity.JwtUser)
	if ok {
		return user.ID, nil
	}

	return 0, errors.New("can't convent to id struct")
}

// GetUserInfo 返回用户信息
func GetUserInfo(c *gin.Context) (*entity.JwtUser, error) {
	u, exist := c.Get(middleware.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get user's info")
	}

	user, ok := u.(*entity.JwtUser)
	if ok {
		return user, nil
	}

	return nil, errors.New("can't convent to userInfo struct")
}
