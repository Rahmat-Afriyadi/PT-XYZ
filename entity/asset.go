package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Asset struct {
	ID         string   `form:"id" json:"id" gorm:"type:varchar(36);primary_key;column:id"`
	NamaAsset  string   `form:"nama_asset" json:"nama_asset" gorm:"type:varchar(100);column:nama_asset"`
	Otr        float64   `form:"otr" json:"otr" gorm:"column:otr"`
	PlatformId string   `form:"platform_id" json:"platform_id" gorm:"type:varchar(36);column:platform_id"`
	Platform   Platform `form:"platform" json:"platform" gorm:"->;references:ID;foreignKey:PlatformId"`
}

func (Asset) TableName() string {
	return "asset"
}

func (e *Asset) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	return
}
