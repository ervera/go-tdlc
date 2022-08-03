package login

import (
	"context"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(ctx context.Context, email string, password string) (domain.User, bool)
}

type service struct {
	userRepository user.Repository
}

func (s *service) CreateAccount() {
}

func (s *service) Login(ctx context.Context, email string, password string) (domain.User, bool) {
	usu, encontrado, _ := s.userRepository.Exists(ctx, email)
	if !encontrado {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	usu.Password = ""
	if err != nil {
		return usu, false
	}
	return usu, true
}

func NewService(r user.Repository) Service {
	return &service{
		userRepository: r,
	}
}
