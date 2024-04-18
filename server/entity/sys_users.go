package entity

import "go_admin/global"

// SysUsers 用户模型
type SysUsers struct {
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

func (u SysUsers) TableName() string {
	return "sys_users"
}

// JwtUser jwt 结构体
type JwtUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

// LoginDto 登录参数
type LoginDto struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Captcha   string `json:"captcha" validate:"required,min=4,max=6"`
	ImageCode string `json:"imageCode" validate:"required"`
}

// RegisterDto 注册参数
type RegisterDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Status   int    `gorm:"column:status;comment:是否启用 1->启用，0->禁用" json:"status"`
}

// VerifyRegisterDto 注册校验参数
type VerifyRegisterDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
