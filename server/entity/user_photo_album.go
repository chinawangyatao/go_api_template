package entity

import "go_admin/global"

type UserPhotoAlbum struct {
	global.GModel
}

func (UserPhotoAlbum) TableName() string { return "user_photo_album" }
