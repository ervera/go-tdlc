package user

import (
	"context"
	"errors"
	"strings"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/iso8601"
	"github.com/ervera/tdlc-gin/pkg/jwt"
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetById(ctx context.Context, guid uuid.UUID) (domain.User, error)
	Update(ctx context.Context, u domain.User) error
	// SaveUserRelation(ctx context.Context, userRelationId string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if len(user.Email) < 3 {
		return domain.User{}, errors.New("email invalido")
	}

	if len(user.Email) < 6 {
		return domain.User{}, errors.New("password debe tener al menos 5 caracteres")
	}

	user.CreatedOn = iso8601.NowTime()

	user, err := s.repository.Save(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return domain.User{}, errors.New("email ingresado ya existe")
		}
		return domain.User{}, errors.New("Ocurrio un error al registrar un usuario: " + err.Error())
	}
	return user, nil
}

func (s *service) GetById(ctx context.Context, ID uuid.UUID) (domain.User, error) {
	user, err := s.repository.GetById(ctx, ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) Update(ctx context.Context, u domain.User) error {
	if jwt.UserID != u.ID {
		return errors.New("userId and token not match")
	}
	return s.repository.Update(ctx, u)
}

// func (s *service) SaveUserRelation(ctx context.Context, userRelationId string) error {
// 	userRelation := domain.UserRelation{UserID: jwt.UserID, UserRelationId: userRelationId}

// 	user, err := s.repository.GetOne(ctx, userRelationId)
// 	if err != nil {
// 		return err
// 	}
// 	if user.Email == "" {
// 		return errors.New("el usuario ingresado no existe")
// 	}
// 	return s.repository.SaveUserRelation(ctx, userRelation)
// }
