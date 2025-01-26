package transaksi

import (
	"errors"
	"fmt"
	"sync"
	"xyz/entity"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	CreateTransaksi(data TransaksiRequest) error
	MasterDataAll() []entity.Transaksi
	DetailTransaksi(nik string) entity.Transaksi
	Update(body entity.Transaksi) error
}

type transaksiRepository struct {
	conn *gorm.DB
}

func NewTransaksiRepository(conn *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		conn: conn,
	}
}

func (lR *transaksiRepository) DetailTransaksi(nik string) entity.Transaksi {
	transaksi := entity.Transaksi{NoKontrak: nik}
	query := lR.conn.Preload("Limit")
	query.Find(&transaksi)
	return transaksi
}

func (lR *transaksiRepository) CreateTransaksi(data TransaksiRequest) error {
	consumer := entity.Konsumen{Nik: data.Nik}
	var wg sync.WaitGroup
	errorRes := make(chan error)
	for _, txn := range data.Transaksi {
		wg.Add(1)
		go func(transaksi entity.Transaksi) {
			defer wg.Done()
			asset := entity.Asset{ID: txn.AssetId}
			lR.conn.Preload("Platform").Find(&asset)
			if asset.NamaAsset == "" {
				errorRes <- errors.New("asset is not found")
			}
			err := lR.conn.Transaction(func(tx *gorm.DB) error {
				// Ambil consumer dari database
				var c entity.Konsumen
				if err := tx.Preload("Limits").First(&c, consumer.Nik).Error; err != nil {
					return fmt.Errorf("failed to fetch consumer: %w", err)
				}

				// Apply loan
				return c.ApplyLoan(tx, transaksi.Tenor, asset.Otr)
			})

			if err != nil {
				errorRes <- err

			} else {
				transaksi.KonsumenNik = consumer.Nik
				transaksi.NamaAsset = asset.NamaAsset
				transaksi.AdminFee = asset.Platform.AdminFee
				transaksi.JumlahBunga = asset.Platform.Bunga
				transaksi.Otr = float64(asset.Otr)
				lR.conn.Save(&transaksi)
				fmt.Printf("Transaction successful: Tenor %d, Amount %.2f\n", transaksi.Tenor, transaksi.Otr)
			}
		}( txn)
	}
	if <-errorRes != nil {
		return <-errorRes
	}
	return nil

}

func (lR *transaksiRepository) Update(data entity.Transaksi) error {
	record := entity.Transaksi{NoKontrak: data.NoKontrak}
	lR.conn.First(&record)
	if record.NamaAsset == "" {
		return errors.New("maaf data tidak ada")
	}
	return nil
}


func (lR *transaksiRepository) MasterDataAll() []entity.Transaksi {
	var transaksi []entity.Transaksi
	lR.conn.Select("id, nama").Find(&transaksi)
	return transaksi
}
