package global

import (
	"time"

	"gorm.io/gorm"
)

type GModel struct {
	ID        uint           `gorm:"primarykey;comment:主键ID" json:"id"`            // 主键ID
	CreatedAt time.Time      `gorm:"type:dateTime(0);autoCreateTime;comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `gorm:"type:dateTime(0);autoUpdateTime;comment:更新时间"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"-"`                  // 删除时间
}
