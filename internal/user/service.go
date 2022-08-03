package user

import (
	"context"
	"errors"

	"github.com/ervera/tdlc-gin/internal/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

type service struct {
	repository Repository
}

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if len(user.Email) < 3 {
		return domain.User{}, errors.New("email invalido")
	}

	if len(user.Email) < 5 {
		return domain.User{}, errors.New("password debe tener al menos 6 caracteres")
	}

	_, encontrado, _ := s.repository.Exists(ctx, user.Email)
	if encontrado {
		return domain.User{}, errors.New("el mail ingresado ya existe")
	}

	_, status, err := s.repository.Save(ctx, user)
	if err != nil {
		return domain.User{}, errors.New("Ocurrio un error al registrar un usuario" + err.Error())
	}

	if !status {
		return domain.User{}, errors.New("No se ha logrado registrar el usuario" + err.Error())
	}

	return user, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
