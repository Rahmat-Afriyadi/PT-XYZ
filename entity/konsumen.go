package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Konsumen struct {
	Nik           string                `form:"nik" json:"nik" gorm:"type:varchar(16);primary_key;column:nik"`
	Fullname         string                `form:"full_name" json:"full_name" gorm:"type:varchar(100);column:full_name"`
	LegalName       string                `form:"legal_name" json:"legal_name" gorm:"column:legal_name"`
	TempatLahir    string               `form:"tempat_lahir" json:"tempat_lahir" gorm:"type:varchar(100);column:tempat_lahir"`
	TglLahir    string             `form:"tgl_lahir" json:"tgl_lahir" gorm:"type:DATE;column:tgl_lahir"`
	FotoKtp          string                `form:"foto_ktp" json:"foto_ktp" gorm:"type:varchar(100);column:foto_ktp"`
	FotoSelfie      string                `form:"foto_selfie" json:"foto_selfie" gorm:"type:varchar(100);column:foto_selfie"`
	Limits       []Limit       `form:"limit" json:"limit" gorm:"foreignKey:KonsumenNik"`
	CreatedBy    string         `form:"created_by" json:"created_by" gorm:"column:created_by"`
	UpdatedBy    string         `form:"updated_by" json:"updated_by" gorm:"column:updated_by"`
	CreatedAt    *time.Time     `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time     `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (e *Konsumen) TableName() string {
	return "konsumen"
}

func (c *Konsumen) ApplyLoan(tx *gorm.DB, tenor uint8, amount float64) error {
	var limit Limit

	// Ambil data limit berdasarkan tenor dengan lock untuk menghindari race condition
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("konsumen_nik = ? AND tenor = ?", c.Nik, tenor).First(&limit).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("limit not found for tenor %d", tenor)
		}
		return fmt.Errorf("failed to fetch limit: %w", err)
	}

	// Cek apakah limit mencukupi
	if limit.Limit < amount {
		return fmt.Errorf("insufficient limit for tenor %d", tenor)
	}

	// Kurangi limit
	limit.Limit -= amount
	if err := tx.Save(&limit).Error; err != nil {
		return fmt.Errorf("failed to update limit: %w", err)
	}

	return nil
}

