package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Platform struct {
	ID    string `form:"id" json:"id" gorm:"type:varchar(36);primary_key;column:id"`
	NamaPlatform string `form:"nama_platform" json:"nama_platform" gorm:"type:varchar(100);column:nama_platform"`
	Bunga        uint8  `form:"jumlah_bunga" json:"jumlah_bunga" gorm:"column:jumlah_bunga"`
	AdminFee     uint64 `form:"admin_fee" json:"admin_fee" gorm:"column:admin_fee"`
	CreatedBy    string         `form:"created_by" json:"created_by" gorm:"column:created_by"`
	UpdatedBy    string         `form:"updated_by" json:"updated_by" gorm:"column:updated_by"`
	CreatedAt    *time.Time     `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time     `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
func (Platform) TableName() string {
	return "platform"
}

func (e *Platform) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}
