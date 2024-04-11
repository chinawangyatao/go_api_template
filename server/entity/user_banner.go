package entity

import "go_admin/global"

// UserBanner 用户banner
type UserBanner struct {
	global.GModel
	Images []string `gorm:"column:images;comment:banner列表" json:"images"`
	UserID uint     `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User   SysUsers `gorm:"foreignKey:UserID" json:"user"`
}

func (u UserBanner) TableName() string {
	return "user_banner"
}
