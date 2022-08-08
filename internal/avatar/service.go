package avatar

import (
	"context"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/jwt"
)

type Service interface {
	UploadAvatar(ctx context.Context, user domain.User) error
}

type service struct {
	userRepository user.Repository
}

func NewService(userRepo user.Repository) Service {
	return &service{
		userRepository: userRepo,
	}
}

func (r *service) UploadAvatar(ctx context.Context, user domain.User) error {
	return r.userRepository.UpdateSelf(ctx, user, jwt.UserID)
}
