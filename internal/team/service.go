package team

import (
	"context"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/iso8601"
	"github.com/ervera/tdlc-gin/pkg/jwt"
)

type Service interface {
	Save(ctx context.Context, team domain.Team) (domain.Team, error)
	GetAll(ctx context.Context) ([]domain.Team, error)
	GetAllByUserId(ctx context.Context) ([]domain.Team, error)
}

type service struct {
	repository Repository
}

func (s *service) Save(ctx context.Context, team domain.Team) (domain.Team, error) {
	team.CreatedOn = iso8601.NowTime()
	team.Enable = true
	return s.repository.Save(ctx, team)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Team, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetAllByUserId(ctx context.Context) ([]domain.Team, error) {
	return s.repository.GetAllByUserId(ctx, jwt.UserID)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
