package usecase

import "context"

func (us *usecase) Logout(ctx context.Context, token string) error {
	jwtData, err := us.jwt.VerifyToken(token)
	if err != nil {
		return err
	}

	return us.userRepo.LogoutUser(ctx, token, jwtData.ExpiresAt.Time)
}
