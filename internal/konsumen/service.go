package konsumen

import "xyz/entity"

type KonsumenService interface {
	CreateKonsumen(data entity.Konsumen) error
	MasterDataAll() []entity.Konsumen
	DetailKonsumen(nik string) entity.Konsumen
	Update(body entity.Konsumen) error
	updateKtp(data entity.Konsumen) error
	updateSelfie(data entity.Konsumen) error
}

type konsumenService struct {
	trR KonsumenRepository
}

func NewKonsumenService(tR KonsumenRepository) KonsumenService {
	return &konsumenService{
		trR: tR,
	}
}

func (s *konsumenService) Update(body entity.Konsumen) error {
	return s.trR.Update(body)
}
func (s *konsumenService) DetailKonsumen(nik string) entity.Konsumen {
	return s.trR.DetailKonsumen(nik)
}

func (s *konsumenService) MasterDataAll() []entity.Konsumen {
	return s.trR.MasterDataAll()
}

func (s *konsumenService) CreateKonsumen(data entity.Konsumen) error {
	return s.trR.CreateKonsumen(data)
}
func (s *konsumenService) updateKtp(data entity.Konsumen) error {
	return s.trR.updateKtp(data)
}
func (s *konsumenService) updateSelfie(data entity.Konsumen) error {
	return s.trR.updateSelfie(data)
}
