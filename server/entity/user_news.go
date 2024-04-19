package entity

import (
	"go_admin/global"
)

type UserNews struct {
	global.GModel
	OpenID       uint     `gorm:"column:open_ID;comment:用户ID" json:"openId"`
	User         SysUsers `gorm:"foreignKey:OpenID" json:"user"`
	CoverImage   string   `gorm:"column:case_title;comment:封面照片" json:"coverImage"`
	NewsTitle    string   `gorm:"column:news_title;comment:新闻标题" json:"newsTitle"`
	NewsSubtitle string   `gorm:"column:news_subtitle;comment:新闻副标题" json:"newsSubtitle"`
	NewsDetails  string   `gorm:"column:news_details;comment:案例详情" json:"newsDetails"`
}

func (UserNews) TableName() string { return "user_news" }
