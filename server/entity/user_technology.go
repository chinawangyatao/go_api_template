package entity

import "go_admin/global"

type UserTechnology struct {
	global.GModel
	UserID     uint     `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User       SysUsers `gorm:"foreignKey:UserID" json:"user"`
	CoverImage string   `gorm:"column:case_title;comment:封面照片" json:"cover_image"`
	Title      string   `gorm:"column:title;comment:标题" json:"title"`
	Details    string   `gorm:"column:details;comment:技术详情" json:"details"`
}

func (UserTechnology) TableName() string {
	return "user_technology"
}
