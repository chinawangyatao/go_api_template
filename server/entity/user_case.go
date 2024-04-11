package entity

import (
	"go_admin/global"
)

type UserCase struct {
	global.GModel
	UserID      uint     `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User        SysUsers `gorm:"foreignKey:UserID" json:"user"`
	Type        string   `gorm:"column:type;comment:案例类型" json:"type"`
	CoverImage  string   `gorm:"column:case_title;comment:封面照片" json:"cover_image"`
	CaseTitle   string   `gorm:"column:case_title;comment:案例标题" json:"case_title"`
	CaseDetails string   `gorm:"column:case_details;comment:案例详情" json:"case_details"`
}

func (UserCase) TableName() string {
	return "user_case"
}
