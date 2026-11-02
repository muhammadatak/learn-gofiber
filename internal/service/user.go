package service

import (
	"context"
	"errors"

	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Get implements [domain.UserService].
func (u *userService) Get(ctx context.Context, email string) (dto.UserData, error) {
	user, err := u.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return dto.UserData{}, err
	}

	if user.Email == "" {
		return dto.UserData{}, errors.New("Data user tidak ditemukan")
	}

	return dto.UserData{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
