package use_cases

import (
	"dtonetest/internal/services"
	"dtonetest/repositories"
	"errors"
)

type LoginDto struct {
	Username string
	Password string
}

type LoginResponse struct {
	Jwt string
}

type ILoginUseCase interface {
	Execute(loginDto LoginDto) (*LoginResponse, error)
}

type LoginUseCase struct {
	userRepo        repositories.UserRepository
	webTokenService services.IWebTokenService
}

func NewLoginUseCase(userRepo repositories.UserRepository, webTokenService services.IWebTokenService) (*LoginUseCase, error) {
	return &LoginUseCase{userRepo: userRepo, webTokenService: webTokenService}, nil
}

func (l LoginUseCase) Execute(in LoginDto) (*LoginResponse, error) {
	user, err := l.userRepo.FindByUsername(in.Username)
	if err != nil {
		return nil, errors.New("the user does not exists")
	}
	valid, err := user.CheckPassword(in.Password)
	if err != nil || !valid {
		return nil, errors.New("the password is incorrect")
	}
	wt, err := l.webTokenService.GenerateToken(in.Username)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		Jwt: wt,
	}, nil
}
