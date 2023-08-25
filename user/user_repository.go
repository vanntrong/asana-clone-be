package user

import (
	"github.com/vanntrong/asana-clone-be/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByEmail(email string) (user *entities.User, err error)
	FindById(id string) (user *entities.User, err error)
	CreateUser(email string, password string, name string) (user *entities.User, err error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) FindByEmail(email string) (user *entities.User, err error) {
	user = &entities.User{}
	result := repo.db.Where("email = ?", email).First(user)
	err = result.Error

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return user, err

}

func (repo *UserRepository) CreateUser(email string, password string, name string) (user *entities.User, err error) {
	user = &entities.User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	err = repo.db.Create(&user).Error

	user.HashPassword()
	return user, err
}

func (repo *UserRepository) FindById(id string) (user *entities.User, err error) {
	user = &entities.User{}
	result := repo.db.Where("id = ?", id).First(user)
	err = result.Error

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return user, err
}
