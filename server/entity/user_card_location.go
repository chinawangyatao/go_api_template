package entity

import "go_admin/global"

// UserCardLocation 用户的位置
type UserCardLocation struct {
	global.GModel
	UserCardID int      `gorm:"column:user_card_id;comment:用户的 cardID" json:"user_card_id;"`
	UserCard   UserCard `gorm:"foreignKey:UserCardID" json:"user_card"`
	Latitude   float64  `gorm:"column:latitude;comment:经度" json:"latitude"`
	Longitude  float64  `gorm:"column:longitude;comment:纬度" json:"longitude"`
}

func (u UserCardLocation) TableName() string {
	return "user_card_location"
}
