package domain

import (
	"context"
	"gofiber-rest-api/dto"
)

type AuthService interface {
	Login(ctx context.Context, request dto.AuthRequest) (dto.AuthResponse, error)
}
