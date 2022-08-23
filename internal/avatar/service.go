package avatar

import (
	"github.com/ervera/tdlc-gin/internal/user"
)

type Service interface {
	//UploadAvatar(ctx context.Context, user domain.User) error
}

type service struct {
	userRepository user.Repository
}

func NewService(userRepo user.Repository) Service {
	return &service{
		userRepository: userRepo,
	}
}

// func (r *service) UploadAvatar(ctx context.Context, user domain.User) error {
// 	return r.userRepository.UpdateSelf(ctx, user, jwt.UserID)
// }
