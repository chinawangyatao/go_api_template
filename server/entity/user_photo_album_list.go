package entity

import "gorm.io/gorm"

// UserPhotoAlbumList 用户相册中间表
type UserPhotoAlbumList struct {
	gorm.Model
	UserPhotoAlbumID int            `gorm:"column:user_photo_album_id;commit:图片" json:"user_photo_album_id"`
	UserPhotoAlbum   UserPhotoAlbum `gorm:"foreignkey:UserPhotoAlbumID" json:"user_photo_album"`
	UserID           uint           `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User             SysUsers       `gorm:"foreignKey:UserID" json:"user"`

	CoverImage string `gorm:"column:case_title;comment:封面照片" json:"cover_image"`
	Title      string `gorm:"column:title;comment:标题" json:"title"`
	Status     int    `gorm:"column:status;comment:是否启用 1->启用，0->禁用" json:"status"`
}

func (u UserPhotoAlbumList) TableName() string {
	return "user_photo_album_list"
}
