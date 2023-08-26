package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vanntrong/asana-clone-be/configs"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenToken(payload map[string]string) (token *Token) {
	var access_token jwt.Token
	var refresh_token jwt.Token

	access_token = *jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "asana-clone",
		"sub":   "access_token",
		"email": payload["email"],
		"name":  payload["name"],
		"id":    payload["id"],
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Minute * 500).Unix(),
	})
	refresh_token = *jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "asana-clone",
		"sub":   "refresh_token",
		"email": payload["email"],
		"name":  payload["name"],
		"id":    payload["id"],
		"iat":   time.Now().Unix(),
		"exp":   time.Now().AddDate(0, 0, 30).Unix(),
	})

	access_token_str, _ := access_token.SignedString([]byte(configs.AppConfig.AccessTokenSecret))
	refresh_token_str, _ := refresh_token.SignedString([]byte(configs.AppConfig.RefreshTokenSecret))

	token = &Token{
		AccessToken:  access_token_str,
		RefreshToken: refresh_token_str,
	}

	return token
}

func ValidateToken(token_str string) (map[string]any, error) {
	token, err := jwt.Parse(token_str, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(configs.AppConfig.AccessTokenSecret), nil
	})

	if !token.Valid || err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

func GetUnixTimeNow() int {
	return int(time.Now().Unix())
}
