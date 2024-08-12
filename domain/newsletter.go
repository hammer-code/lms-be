package domain

import (
	"context"
	"net/http"
	"time"
)

type NewsletterRepository interface {
	Subscribe(context.Context, Newsletter) error
	GetByEmail(context.Context, string) (*Newsletter, error)
}

type NewslettersUsecase interface {
	Subscribe(context.Context, string) error
}

type NewslettterHandler interface {
	Subscribe(w http.ResponseWriter, r *http.Request)
}
type Newsletter struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Newsletter) TableName() string {
	return "newsletters"
}

type SubscribeReq struct {
	Email string `json:"email" binding:"required"`
}
