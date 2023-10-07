package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/vanntrong/asana-clone-be/configs"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/user"
	"github.com/vanntrong/asana-clone-be/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type IAuthService interface {
	Register(payload RegisterValidation) (*entities.User, *utils.Token, error)
	Login(payload LoginValidation) (*entities.User, *utils.Token, error)
	LoginGoogle(payload LoginGoogleValidation) (*entities.User, *utils.Token, error)
	CheckEmail(payload CheckEmailValidation) (*entities.User, error)
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
		return &entities.User{}, nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return &entities.User{}, nil, err
	}

	exitsUser, err = service.userRepository.CreateUser(payload.Email, string(hashedPassword), payload.Name)

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

func (service *AuthService) Login(payload LoginValidation) (*entities.User, *utils.Token, error) {
	var exitsUser, err = service.userRepository.FindByEmail(payload.Email)

	if err != nil {
		return &entities.User{}, nil, err
	}

	if exitsUser == nil {
		return &entities.User{}, nil, errors.New("email not exists")
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

func (service *AuthService) LoginGoogle(data LoginGoogleValidation) (*entities.User, *utils.Token, error) {
	ctx := context.Background()
	audience := configs.AppConfig.GoogleClientId
	payload, err := idtoken.Validate(ctx, data.IdToken, audience)

	if err != nil {
		fmt.Print(err)
		return &entities.User{}, nil, nil

	}
	fmt.Print(payload.Claims)

	email := payload.Claims["email"].(string)

	exitsUser, _ := service.userRepository.FindByEmail(email)

	if exitsUser == nil {
		dataCreateUser := user.CreateUserValidation{
			Email:    email,
			Name:     payload.Claims["name"].(string),
			Avatar:   payload.Claims["picture"].(string),
			Provider: "google",
			IsActive: true,
		}
		exitsUser, err = service.userRepository.CreateUserWithProvider(dataCreateUser)
		if err != nil {
			return &entities.User{}, nil, err
		}
	}

	token := utils.GenToken(map[string]string{
		"email": exitsUser.Email,
		"name":  exitsUser.Name,
		"id":    exitsUser.ID.String(),
	})

	return exitsUser, token, nil
}

func (service *AuthService) CheckEmail(payload CheckEmailValidation) (*entities.User, error) {
	var exitsUser, err = service.userRepository.FindByEmail(payload.Email)

	if err != nil {
		return &entities.User{}, err
	}

	if exitsUser == nil {
		return &entities.User{}, errors.New("email not exists")
	}

	return exitsUser, nil
}
