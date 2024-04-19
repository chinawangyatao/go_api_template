package entity

import "go_admin/global"

type UserTechnology struct {
	global.GModel
	OpenID     uint     `gorm:"column:open_ID;comment:用户ID" json:"openId"`
	User       SysUsers `gorm:"foreignKey:OpenID" json:"user"`
	CoverImage string   `gorm:"column:case_title;comment:封面照片" json:"coverImage"`
	Title      string   `gorm:"column:title;comment:标题" json:"title"`
	TecDetails string   `gorm:"column:details;comment:技术详情" json:"TecDetails"`
}

func (UserTechnology) TableName() string {
	return "user_technology"
}
