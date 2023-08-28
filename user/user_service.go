package user

import (
	"github.com/vanntrong/asana-clone-be/common"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/utils"
)

type IUserService interface {
	GetById(id string) (*entities.User, error)
	GetList(query GetListUserQuery) ([]*entities.User, *common.PaginationResponse, error)
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

func (service *UserService) GetList(query GetListUserQuery) ([]*entities.User, *common.PaginationResponse, error) {
	users, total, err := service.userRepository.GetList(query)

	if err != nil {
		return nil, nil, err
	}

	pagination := utils.GetPaginationResponse(total, &query.Pagination)

	return users, pagination, nil
}
