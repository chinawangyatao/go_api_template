package entity

import "go_admin/global"

type UserPhotoAlbum struct {
	global.GModel
	UserID uint         `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User   SysUsers     `gorm:"foreignKey:UserID" json:"user"`
	Images []PhotoAlbum `gorm:"column:images;comment:图片组" json:"images"`
}

func (UserPhotoAlbum) TableName() string { return "user_photo_album" }

type PhotoAlbum struct {
	CoverImage []string `gorm:"column:case_title;comment:封面照片" json:"cover_image"`
	Title      string   `gorm:"column:title;comment:标题" json:"title"`
}
