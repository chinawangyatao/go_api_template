package entity

import (
	"go_admin/global"
)

type UserNews struct {
	global.GModel
	UserID       uint     `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User         SysUsers `gorm:"foreignKey:UserID" json:"user"`
	CoverImage   string   `gorm:"column:case_title;comment:封面照片" json:"cover_image"`
	NewsTitle    string   `gorm:"column:news_title;comment:新闻标题" json:"news_title"`
	NewsSubtitle string   `gorm:"column:news_subtitle;comment:新闻副标题" json:"news_subtitle"`
	Details      string   `gorm:"column:details;comment:案例详情" json:"details"`
}

func (UserNews) TableName() string { return "user_news" }
