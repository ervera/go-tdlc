package tweet

import (
	"context"
	"time"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/jwt"
)

type Service interface {
	Save(ctx context.Context, message string) (domain.Tweet, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(ctx context.Context, message string) (domain.Tweet, error) {
	tweet := domain.Tweet{
		Mensaje: message,
		UserID:  jwt.UserID,
		Fecha:   time.Now(),
	}

	result, err := s.repository.Save(ctx, tweet)
	return result, err
}
