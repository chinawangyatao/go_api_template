package entity

import "gorm.io/gorm"

// UserBannerList banner 中间表
type UserBannerList struct {
	gorm.Model
	UserBannerID uint       `gorm:"commit:轮播的 id" json:"userBannerId"`
	UserBanner   UserBanner `gorm:"foreignKey:UserBannerID" json:"userBanner"`
	UserID       uint       `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	User         SysUsers   `gorm:"foreignKey:UserID" json:"user"`

	Status int    `gorm:"column:image;comment:是否启用 1->启用，0->禁用" json:"status"`
	Image  string `gorm:"column:image; type:varchar(255);comment:图片地址" json:"image"`
}

func (u UserBannerList) TableName() string {
	return "user_banner_list"
}
