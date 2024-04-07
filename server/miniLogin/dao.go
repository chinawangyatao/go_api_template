package miniLogin

import (
	"go_admin/pkg/db"
	"go_admin/server/entity"
	"go_admin/utils"
)

// MiniUserInfo 小程序用户查询，如果有就返回用户信息，如果没有就落库
func MiniUserInfo(dto entity.SysUsers) (miniUserInfo entity.SysUsers) {
	openId := dto.OpenId
	err := db.DB.Where("openid=?", openId).First(&miniUserInfo)
	if err != nil { // 如果没有落库
		utils.L.Info("where openid err:", err)
		db.DB.Create(&dto)
		return dto
	}
	return miniUserInfo
}
