package konsumen

import (
	"errors"
	"fmt"
	"xyz/entity"

	"gorm.io/gorm"
)

type KonsumenRepository interface {
	CreateKonsumen(data entity.Konsumen) error
	MasterDataAll() []entity.Konsumen
	DetailKonsumen(nik string) entity.Konsumen
	Update(body entity.Konsumen) error
	updateKtp(data entity.Konsumen) error
	updateSelfie(data entity.Konsumen) error
}

type konsumenRepository struct {
	conn *gorm.DB
}

func NewKonsumenRepository(conn *gorm.DB) KonsumenRepository {
	return &konsumenRepository{
		conn: conn,
	}
}

func (lR *konsumenRepository) DetailKonsumen(nik string) entity.Konsumen {
	konsumen := entity.Konsumen{Nik: nik}
	query := lR.conn.Preload("Limits")
	query.Find(&konsumen)
	return konsumen
}

func (lR *konsumenRepository) CreateKonsumen(data entity.Konsumen) error {
	limit := data.Limits
	result := lR.conn.Omit("Limits").Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	}
	err := lR.conn.Model(&data).Association("Limits").Replace(limit)
	if err != nil {
		return err
	}
	return nil

}

func (lR *konsumenRepository) Update(data entity.Konsumen) error {
	record := entity.Konsumen{Nik: data.Nik}
	lR.conn.First(&record)
	if record.Fullname == "" {
		return errors.New("maaf data tidak ada")
	}
	limit := data.Limits

	err := lR.conn.Model(&record).Association("Limits").Replace(limit)
	if err != nil {
		return err
	}
	return nil
}

func (lR *konsumenRepository) updateKtp(data entity.Konsumen) error {
	record := entity.Konsumen{Nik: data.Nik}
	lR.conn.First(&record)
	if record.Fullname == "" {
		return errors.New("maaf data tidak ada")
	}
	record.FotoKtp = data.FotoKtp
	lR.conn.Save(&record)
	return nil
}

func (lR *konsumenRepository) updateSelfie(data entity.Konsumen) error {
	record := entity.Konsumen{Nik: data.Nik}
	lR.conn.First(&record)
	if record.Fullname == "" {
		return errors.New("maaf data tidak ada")
	}
	record.FotoSelfie = data.FotoSelfie
	lR.conn.Save(&record)
	return nil
}


func (lR *konsumenRepository) MasterDataAll() []entity.Konsumen {
	var konsumen []entity.Konsumen
	lR.conn.Find(&konsumen)
	return konsumen
}
