package entity

import "go_admin/global"

// UserPhotoAlbum 用户相册
type UserPhotoAlbum struct {
	global.GModel
	UserID uint     `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User   SysUsers `gorm:"foreignKey:UserID" json:"user"`
}

func (UserPhotoAlbum) TableName() string { return "user_photo_album" }
