package domain

import "time"

type (
	User struct {
		ID        int    `gorm:"primaryKey" json:"id"`
		Username  string `json:"username" gorm:"type:varchar(255);not null"`
		Email     string `json:"email" gorm:"type:varchar(255);not null"`
		Password  string `json:"password" gorm:"type:varchar(255);not null"`
		Role      string  `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Register struct {
		Username        string `json:"username" binding:"required"`
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	Login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

// define gorm table of struct
func (User) TableName() string {
	return "users"
}

func RegistToUser(r Register) User {
	return User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}
