package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/iso8601"
	"github.com/ervera/tdlc-gin/pkg/jwt"
	"github.com/ervera/tdlc-gin/pkg/random"
	"github.com/ervera/tdlc-gin/pkg/sendgrid"
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetById(ctx context.Context, guid uuid.UUID) (domain.User, error)
	Update(ctx context.Context, u domain.User) error
	SendEmailWithPassword(ctx context.Context, email string) error
	NewPassword(ctx context.Context, password string) error
	// SaveUserRelation(ctx context.Context, userRelationId string) error
}

type service struct {
	repository      Repository
	sendgridService sendgrid.Service
}

func NewService(r Repository, s sendgrid.Service) Service {
	return &service{
		repository:      r,
		sendgridService: s,
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

func (s *service) SendEmailWithPassword(ctx context.Context, email string) error {
	user, err := s.repository.ExistAndGetByMail(ctx, email)
	if err != nil {
		return err
	}
	randomPassword := random.GenerateStringByN(6)
	err = s.repository.UpdatePasswordById(ctx, user, randomPassword)
	if err != nil {
		return err
	}
	s.sendgridService.SendPassword(ctx, user.FirstName, email, randomPassword)
	return nil
}

func (s *service) NewPassword(ctx context.Context, password string) error {
	user, err := s.repository.ExistAndGetByMail(ctx, jwt.Email)
	if err != nil {
		fmt.Println(jwt.Email)
		fmt.Println("A")
		return err
	}
	err = s.repository.UpdatePasswordById(ctx, user, password)
	if err != nil {
		fmt.Println("b")
		return err
	}
	return nil
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
