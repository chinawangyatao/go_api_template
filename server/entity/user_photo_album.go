package entity

import "go_admin/global"

// UserPhotoAlbum 用户相册
type UserPhotoAlbum struct {
	global.GModel
	OpenID uint     `gorm:"column:open_ID;comment:用户ID" json:"openId"`
	User   SysUsers `gorm:"foreignKey:OpenID" json:"user"`
}

func (UserPhotoAlbum) TableName() string { return "user_photo_album" }
