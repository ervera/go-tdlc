package login

import (
	"context"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, email string, password string) (domain.User, error)
}

type service struct {
	userRepository user.Repository
}

func (s *service) Login(ctx context.Context, email string, password string) (domain.User, error) {
	usu, err := s.userRepository.ExistAndGetByMail(ctx, email)
	if err != nil {
		return usu, err
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	usu.Password = ""
	if err != nil {
		return usu, err
	}
	return usu, nil
}

func NewService(r user.Repository) Service {
	return &service{
		userRepository: r,
	}
}
