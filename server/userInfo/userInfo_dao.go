package userInfo

import (
	"go_admin/pkg/db"
	"go_admin/server/entity"
	"go_admin/utils"
)

// SysUserInfoDetails 用户详情
func SysUserInfoDetails(dto entity.LoginDto) (sysUserInfo entity.SysUsers) {
	username := dto.Username
	err := db.DB.Where("username =?", username).First(&sysUserInfo)
	if err != nil {
		utils.L.Info("where username err:", err)
	}
	return sysUserInfo
}

// SysUserInfoSave 用户注册落库
func SysUserInfoSave(dto entity.SysUsers) (info entity.SysUsers, error error) {
	info = dto
	err := db.DB.Create(&info)
	if err != nil {
		utils.L.Info("save userInfo err:", err.Error)
		return info, err.Error
	}
	return info, nil
}

// VerifyInfoUsername 校验用户信息
func VerifyInfoUsername(dto entity.VerifyRegisterDto) (userInfo entity.SysUsers) {
	username := dto.Username
	email := dto.Email
	err := db.DB.Where("username=?", username).Or("email=?", email).First(&userInfo)
	if err != nil {
		utils.L.Info("VerifyInfoUsername err:", err)
	}
	return userInfo
}
