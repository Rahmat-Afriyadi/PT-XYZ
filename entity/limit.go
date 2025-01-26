package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Limit struct {
	ID          string   `form:"id" json:"id" gorm:"type:varchar(36);primary_key;column:id"`
	Tenor       uint8    `form:"tenor" json:"tenor" gorm:"column:tenor"`
	Limit       float64   `form:"limit" json:"limit" gorm:"column:limit"`
	KonsumenNik string   `form:"konsumen_nik" json:"konsumen_nik" gorm:"type:varchar(16);column:konsumen_nik"`
	Konsumen    Konsumen `form:"konsumen" json:"konsumen" gorm:"->;references:Nik;foreignKey:KonsumenNik"`
	CreatedBy    string         `form:"created_by" json:"created_by" gorm:"column:created_by"`
	UpdatedBy    string         `form:"updated_by" json:"updated_by" gorm:"column:updated_by"`
	CreatedAt    *time.Time     `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time     `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Limit) TableName() string {
	return "limits"
}

func (e *Limit) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}
