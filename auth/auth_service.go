package auth

import (
	"errors"

	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/user"
	"github.com/vanntrong/asana-clone-be/utils"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(payload RegisterValidation) (*entities.User, *utils.Token, error)
	Login(payload LoginValidation) (*entities.User, *utils.Token, error)
}

type AuthService struct {
	userRepository user.IUserRepository
}

func NewAuthService(userRepository user.IUserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (service *AuthService) Register(payload RegisterValidation) (*entities.User, *utils.Token, error) {
	var exitsUser, err = service.userRepository.FindByEmail(payload.Email)

	if err != nil {
		return &entities.User{}, nil, err
	}

	if exitsUser != nil {
		return &entities.User{}, nil, errors.New("Email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return &entities.User{}, nil, err
	}

	exitsUser, err = service.userRepository.CreateUser(payload.Email, string(hashedPassword), payload.Name)

	token := utils.GenToken(map[string]string{
		"email": exitsUser.Email,
		"name":  exitsUser.Name,
		"id":    exitsUser.ID.String(),
	})

	return exitsUser, token, nil
}

func (service *AuthService) Login(payload LoginValidation) (*entities.User, *utils.Token, error) {
	var exitsUser, err = service.userRepository.FindByEmail(payload.Email)

	if err != nil {
		return &entities.User{}, nil, err
	}

	if exitsUser == nil {
		return &entities.User{}, nil, errors.New("Email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(exitsUser.Password), []byte(payload.Password))

	if err != nil {
		return &entities.User{}, nil, err
	}

	token := utils.GenToken(map[string]string{
		"email": exitsUser.Email,
		"name":  exitsUser.Name,
		"id":    exitsUser.ID.String(),
	})

	return exitsUser, token, nil
}
