package usecase

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
)

func (us *usecase) GetUsers(ctx context.Context) (users []domain.User, err error) {
	return us.userRepo.GetUsers(ctx)
}
