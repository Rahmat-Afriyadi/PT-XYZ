package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	NoKontrak   string   `form:"no_kontrak" json:"no_kontrak" gorm:"type:varchar(16);primary_key;column:no_kontrak"`
	NamaAsset   string   `form:"nama_asset" json:"nama_asset" gorm:"type:varchar(100);column:nama_asset"`
	Tenor       uint8    `form:"jumlah_cicilan" json:"jumlah_cicilan" gorm:"column:jumlah_cicilan"`
	JumlahBunga uint8    `form:"jumlah_bunga" json:"jumlah_bunga" gorm:"column:jumlah_bunga"`
	Otr         float64   `form:"otr" json:"otr" gorm:"column:otr"`
	AdminFee    uint64   `form:"admin_fee" json:"admin_fee" gorm:"column:admin_fee"`
	KonsumenNik string   `form:"konsumen_nik" json:"konsumen_nik" gorm:"type:varchar(16);column:konsumen_nik"`
	Konsumen    Konsumen `form:"konsumen" json:"konsumen" gorm:"->;references:Nik;foreignKey:KonsumenNik"`
	AssetId     string   `form:"asset_id" json:"asset_id" gorm:"type:varchar(36);column:asset_id"`
	Asset       Asset    `form:"asset" json:"asset" gorm:"->;references:ID;foreignKey:AssetId"`
	CreatedBy    string         `form:"created_by" json:"created_by" gorm:"column:created_by"`
	UpdatedBy    string         `form:"updated_by" json:"updated_by" gorm:"column:updated_by"`
	CreatedAt    *time.Time     `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time     `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
func (Transaksi) TableName() string {
	return "transaksi"
}

func (u *Transaksi) BeforeCreate(tx *gorm.DB) (err error) {
	if u.NoKontrak != "" {
		return
	}
	lastTransaksi := Transaksi{}
	tx.Last(&lastTransaksi)
	if lastTransaksi.NoKontrak != "" {
		u.NoKontrak = GenerateNomorKontrak(lastTransaksi)
	} else {
		u.NoKontrak = "TRX-001"
	}
	return
}


func GenerateNomorKontrak(transaksi Transaksi) string {

	splitId := strings.Split(transaksi.NoKontrak, "-")
	i, err := strconv.Atoi(splitId[1])
	if err != nil {
		fmt.Println("ini error parse string to int ", err)
	}
	i += 1
	idProduk := ""
	if i > 99 {
		idProduk = fmt.Sprintf("%s%d", splitId[0]+"-", i)
	} else if i > 9 {
		idProduk = fmt.Sprintf("%s%d", splitId[0]+"-0", i)
	} else {
		idProduk = fmt.Sprintf("%s%d", splitId[0]+"-00", i)
	}
	return idProduk

}
