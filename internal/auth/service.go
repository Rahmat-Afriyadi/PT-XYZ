package auth

import "xyz/entity"

type AuthService interface {
	SignInUser(r SigninRequest) (entity.User, error)
	SignInUserAsuransi(r SigninRequest) (entity.User, error)
	RefreshToken(r string) (entity.User, error)
	RefreshTokenAsuransi(r string) (entity.User, error)
}

type authService struct {
	uR UserRepository
}

func NewAuthService(uR UserRepository) AuthService {
	return &authService{
		uR,
	}
}

func (s *authService) SignInUser(r SigninRequest) (entity.User, error) {
	return s.uR.FindByUsername(r.Username), nil
}

func (s *authService) SignInUserAsuransi(r SigninRequest) (entity.User, error) {
	return s.uR.FindByUsername(r.Username), nil
}

func (s *authService) RefreshToken(r string) (entity.User, error) {
	return s.uR.FindById(r), nil
}

func (s *authService) RefreshTokenAsuransi(r string) (entity.User, error) {
	return s.uR.FindById(r), nil
}

