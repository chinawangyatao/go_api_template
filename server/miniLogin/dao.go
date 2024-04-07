package miniLogin

import (
	"errors"
	"go_admin/pkg/db"
	"go_admin/server/entity"
	"go_admin/utils"
	"gorm.io/gorm"
)

// MiniUserInfo 小程序用户查询，如果有就返回用户信息，如果没有就落库
func MiniUserInfo(dto entity.SysUsers) (miniUserInfo entity.SysUsers) {
	openId := dto.OpenId
	var err error
	err = db.DB.Where("openid=?", openId).First(&miniUserInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.L.Info("Record not found for openid:", openId)
			// 如果没有找到记录，创建新记录并返回
			if err := db.DB.Create(&dto).Error; err != nil {
				utils.L.Error("Error creating new record:", err)
				return
			}
			return dto
		}
		utils.L.Error("Error querying database:", err)
		return
	}
	return miniUserInfo
}
