package user

import (
	"github.com/vanntrong/asana-clone-be/entities"
)

type IUserService interface {
	GetById(id string) (*entities.User, error)
}

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) GetById(id string) (*entities.User, error) {
	user, err := service.userRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
