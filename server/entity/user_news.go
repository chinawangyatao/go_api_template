package entity

import (
	"go_admin/global"
)

type UserNews struct {
	global.GModel
}

func (UserNews) TableName() string { return "user_news" }
