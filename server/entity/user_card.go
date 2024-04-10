package entity

import "go_admin/global"

type UserCard struct {
	global.GModel
}

func (c UserCard) tableName() string {
	return "user_card"
}
