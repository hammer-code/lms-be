package domain

import "time"

type (
	LogoutToken struct {
		ID        int       `gorm:"primaryKey" json:"id"`
		Token     string    `gorm:"type:varchar(255);not null;unique" json:"token"`
		ExpiredAt time.Time `json:"expired_at"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func (LogoutToken) TableName() string {
	return "logout"
}
