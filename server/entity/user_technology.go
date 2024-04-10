package entity

import "go_admin/global"

type UserTechnology struct {
	global.GModel
}

func (UserTechnology) TableName() string {
	return "user_technology"
}
