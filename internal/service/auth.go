package service

import (
	"context"
	"errors"
	"gofiber-rest-api/domain"
	"gofiber-rest-api/dto"
	"gofiber-rest-api/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuth(conf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return authService{
		conf:           conf,
		userRepository: userRepository,
	}
}

func (a authService) Login(ctx context.Context, request dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	if user.ID == "" {
		return dto.AuthResponse{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("password incorrect")
	}
	claim := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		Token: tokenString,
	}, nil

}
