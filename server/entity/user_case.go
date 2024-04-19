package entity

import (
	"go_admin/global"
)

type UserCase struct {
	global.GModel
	OpenID      uint     `gorm:"column:open_ID;comment:用户ID" json:"openId"`
	User        SysUsers `gorm:"foreignKey:OpenID" json:"user"`
	Type        string   `gorm:"column:type;comment:案例类型" json:"type"`
	CoverImage  string   `gorm:"column:case_title;comment:封面照片" json:"coverImage"`
	CaseTitle   string   `gorm:"column:case_title;comment:案例标题" json:"caseTitle"`
	CaseDetails string   `gorm:"column:case_details;comment:案例详情" json:"caseDetails"`
}

func (UserCase) TableName() string {
	return "user_case"
}
