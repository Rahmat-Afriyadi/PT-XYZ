package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string   `form:"id" json:"id" gorm:"primary_key;column:id"`
	Nik        string   `form:"nik" json:"nik" gorm:"column:nik"`
	Name        string   `form:"name" json:"name" gorm:"column:name"`
	Password    string   `form:"password" json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

