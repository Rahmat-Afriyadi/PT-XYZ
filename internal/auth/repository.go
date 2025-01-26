package auth

import (
	"xyz/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id string) entity.User
	FindByUsername(username string) entity.User
}

type userRepository struct {
	connUser *gorm.DB
}

func NewUserRepository(connUser *gorm.DB) UserRepository {
	return &userRepository{
		connUser: connUser,
	}
}

func (lR *userRepository) FindById(id string) entity.User {
	user := entity.User{ID: id}
	lR.connUser.Find(&user)


	return user
}


func (lR *userRepository) FindByUsername(username string) entity.User {
	var user entity.User
	lR.connUser.Where("nik", username).First(&user)

	return user
}

