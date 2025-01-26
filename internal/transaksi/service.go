package transaksi

import "xyz/entity"

type TransaksiService interface {
	CreateTransaksi(data TransaksiRequest) error
	MasterDataAll() []entity.Transaksi
	DetailTransaksi(nik string) entity.Transaksi
	Update(body entity.Transaksi) error
}

type transaksiService struct {
	trR TransaksiRepository
}

func NewTransaksiService(tR TransaksiRepository) TransaksiService {
	return &transaksiService{
		trR: tR,
	}
}

func (s *transaksiService) Update(body entity.Transaksi) error {
	return s.trR.Update(body)
}
func (s *transaksiService) DetailTransaksi(nik string) entity.Transaksi {
	return s.trR.DetailTransaksi(nik)
}

func (s *transaksiService) MasterDataAll() []entity.Transaksi {
	return s.trR.MasterDataAll()
}

func (s *transaksiService) CreateTransaksi(data TransaksiRequest) error {
	return s.trR.CreateTransaksi(data)
}
