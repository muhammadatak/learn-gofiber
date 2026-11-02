package domain

import (
	"context"

	"github.com/muhammadatak/learn-gofiber/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
