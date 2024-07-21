package use_cases

import (
	"dtonetest/internal/services"
	"dtonetest/repositories"
	"errors"
)

type LoginDto struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Jwt string `json:"jwt"`
}

type ILoginUseCase interface {
	Execute(loginDto LoginDto) (*LoginResponse, error)
}

type LoginUseCase struct {
	userRepo        repositories.UserRepository
	webTokenService services.IWebTokenService
}

func NewLoginUseCase(userRepo repositories.UserRepository, webTokenService services.IWebTokenService) *LoginUseCase {
	return &LoginUseCase{userRepo: userRepo, webTokenService: webTokenService}
}

func (l LoginUseCase) Execute(in LoginDto) (*LoginResponse, error) {
	user, err := l.userRepo.FindByEmail(in.Email)
	if err != nil {
		return nil, errors.New("the user does not exists")
	}
	valid, err := user.CheckPassword(in.Password)
	if err != nil || !valid {
		return nil, errors.New("the password is incorrect")
	}
	wt, err := l.webTokenService.GenerateToken(in.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		Jwt: wt,
	}, nil
}
