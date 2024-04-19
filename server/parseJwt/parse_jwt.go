package parseJwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/pkg/jwt"
	"go_admin/server/entity"
	"go_admin/utils"
)

func ParseMiniJwt(c *gin.Context) (userInfo *entity.JwtUser, err error) {
	authHeader := c.GetHeader("Authorization")
	userInfo, err = jwt.ParseToken(authHeader)
	if err != nil {
		fmt.Println("parse token err:", err)
		utils.L.Error("parse token err:", err.Error())
		utils.Failed(c, 500, "server err:"+err.Error())
		return userInfo, err
	}
	return userInfo, nil
}
