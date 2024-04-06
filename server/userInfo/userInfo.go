package userInfo

import (
	"github.com/gin-gonic/gin"
	"go_admin/server/entity"
	"go_admin/utils"
)

// Login 登录方法
func Login(c *gin.Context) {
	var dto entity.LoginDto
	err := c.BindJSON(&dto)
	if err != nil {
		utils.L.Error("userinfo login bindjson err: ", err)
		return
	}
	SysUserInfoServiceFunc().Login(c, dto)
}

// Register 注册方法
func Register(c *gin.Context) {
	var dto entity.SysUsers
	err := c.BindJSON(&dto)
	if err != nil {
		utils.L.Error("userinfo Register bindjson err: ", err)
		return
	}
	SysUserInfoServiceFunc().Register(c, dto)
}

// VerifyUsername 验证用户名密码是否可用
func VerifyUsername(c *gin.Context) {
	var dto entity.VerifyRegisterDto
	err := c.BindJSON(&dto)
	if err != nil {
		utils.L.Error("userinfo VerifyUsername bindjson err: ", err)
		return
	}
	SysUserInfoServiceFunc().VerifyUsername(c, dto)
}
