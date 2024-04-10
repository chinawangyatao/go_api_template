package entity

import (
	"go_admin/global"
)

type UserCase struct {
	global.GModel
}

func (UserCase) TableName() string {
	return "user_case"
}
