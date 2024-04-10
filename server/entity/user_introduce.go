package entity

import "go_admin/global"

type UserIntroduce struct {
	global.GModel
}

func (i UserIntroduce) tableName() string {
	return "user_introduce"
}
