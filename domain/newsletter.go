package domain

import "time"

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
