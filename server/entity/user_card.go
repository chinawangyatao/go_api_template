package entity

import "go_admin/global"

type UserCard struct {
	global.GModel
	OpenID   uint     `gorm:"column:open_ID;comment:用户ID" json:"openId"`
	User     SysUsers `gorm:"foreignKey:OpenID" json:"user"`
	Title    string   `gorm:"column:title;comment:标题" json:"title"`
	Subtitle string   `gorm:"column:subtitle;comment:副标题" json:"subtitle"`
	Avatar   string   `gorm:"column:avatar;varchar(500);comment:头像" json:"avatar"`
	Email    string   `gorm:"column:email;varchar(64);comment:邮箱" json:"email"`
	Phone    string   `gorm:"column:phone;varchar(64);comment:手机" json:"phone"`
	Nickname string   `gorm:"column:nick_name;varchar(64);comment:昵称" json:"nickname"`
	JobTitle string   `gorm:"column:job_title;varchar(64);comment:称呼" json:"jobTitle"`
	Label    string   `gorm:"column:label;comment:标签" json:"label"`
	WechatQd string   `gorm:"column:wechat_qd;comment:微信二维码" json:"wechat_qd"`
}

func (c UserCard) TableName() string {
	return "user_card"
}
