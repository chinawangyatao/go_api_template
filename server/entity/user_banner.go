package entity

import "go_admin/global"

// UserBanner 用户banner
type UserBanner struct {
	global.GModel
	Images string   `gorm:"column:images;comment:banner列表" json:"images"`
	OpenID uint     `gorm:"column:open_ID;comment:用户ID" json:"openId"`
	User   SysUsers `gorm:"foreignKey:OpenID" json:"user"`
}

func (u UserBanner) TableName() string {
	return "user_banner"
}
