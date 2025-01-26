package transaksi

import "xyz/entity"

type TransaksiRequest struct {
	Nik string `json:"nik"`
	Transaksi []entity.Transaksi `json:"transaksi"`
}