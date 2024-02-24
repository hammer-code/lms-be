package usecase

import "context"

func (us *usecase) Logout(ctx context.Context, token string) error {
	jwtData, err := us.jwt.VerifyToken(token)
	if err != nil {
		return err
	}

	err = us.dbTX.StartTransaction(ctx, func(txCtx context.Context) error {
		return us.userRepo.LogoutUser(txCtx, token, jwtData.ExpiresAt.Time)
	})

	return err
}
