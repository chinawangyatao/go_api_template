package entity

import "go_admin/global"

// UserBanner 用户banner
type UserBanner struct {
	global.GModel
	PostId   int    `gorm:"column:post_id;comment:岗位 ID" json:"postId"`
	DepId    int    `gorm:"column:Dep_id;comment:部门 ID" json:"DepId"`
	Username string `gorm:"column:username;varchar(64);comment:用户名;NOT NULL" json:"username"`
	Password string `gorm:"column:password;varchar(64);comment:密码;NOT NULL" json:"password"`
	Nickname string `gorm:"column:nick_name;varchar(64);comment:昵称" json:"nickname"`
	Status   int    `gorm:"column:status;comment:账号状态：1->启用，0->禁用;NOT NULL;default:1" json:"status"`
	Avatar   string `gorm:"column:avatar;varchar(500);comment:头像" json:"avatar"`
	Email    string `gorm:"column:email;varchar(64);comment:邮箱" json:"email"`
	Phone    string `gorm:"column:phone;varchar(64);comment:手机" json:"phone"`
	Note     string `gorm:"column:note;varchar(500);comment:备注" json:"note"`
	OpenId   string `gorm:"column:openid;varchar(200);comment:微信小程序 openid" json:"openid"`
}

func (u UserBanner) TableName() string {
	return "user_banner"
}
