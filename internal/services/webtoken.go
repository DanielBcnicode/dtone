package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type IWebTokenService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) error
}
type WebTokenService struct {
	apiSecret string
	lifeSpan  int
}

func NewWebTokenService(apiSecret string, lifeSpan int) (*WebTokenService, error) {
	if len(apiSecret) == 0 {
		return nil, errors.New("secret is required")
	}
	if lifeSpan < 1 {
		return nil, errors.New("lifeSpan is required and must be > 0")
	}
	return &WebTokenService{apiSecret: apiSecret, lifeSpan: lifeSpan}, nil
}

func (wt *WebTokenService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(wt.lifeSpan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(wt.apiSecret))
}

func (wt *WebTokenService) ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(wt.apiSecret), nil
	})
	return err
}
