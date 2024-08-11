package newsletters

import (
	"context"

	"github.com/hammer-code/lms-be/domain"
)

type (
	NewsletterRepository interface {
		Subscribe(context.Context, domain.Newsletter) error
		GetByEmail(context.Context, string) (*domain.Newsletter, error)
	}

	NewslettersUsecase interface {
		Subscribe(context.Context, string) error
	}
)
