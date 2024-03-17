package domain

import "time"

type (
	User struct {
		ID          int       `gorm:"primaryKey" json:"id"`
		Username    string    `json:"username" gorm:"type:varchar(255);not null;unique"`
		Email       string    `json:"email" gorm:"type:varchar(255);not null;unique"`
		Password    string    `json:"password" gorm:"type:varchar(255);not null"`
		Role        string    `json:"role"`
		Fullname    string    `json:"fullname"`
		DateOfBirth time.Time `json:"date_of_birth"`
		Gender      string    `json:"gender"`
		PhoneNumber string    `json:"phone_number"`
		Address     string    `json:"address"`
		Github      string    `json:"github"`
		Linkedin    string    `json:"linkedin"`
		PersonalWeb string    `json:"personal_web"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	UserUpdateProfile struct {
		Fullname    string    `json:"fullname" gorm:"type:varchar(255);not null"`
		DateOfBirth time.Time `json:"date_of_birth" gorm:"type:date;not null"`
		Gender      string    `json:"gender" gorm:"type:varchar(255);not null"`
		PhoneNumber string    `json:"phone_number" gorm:"type:varchar(255); not null"`
		Address     string    `json:"address"`
		Github      string    `json:"github"`
		Linkedin    string    `json:"linkedin"`
		PersonalWeb string    `json:"personal_web"`
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

func UserUpdateProfileToUser(u UserUpdateProfile) User {
	return User{

		Fullname:    u.Fullname,
		DateOfBirth: u.DateOfBirth,
		Gender:      u.Gender,
		PhoneNumber: u.PhoneNumber,
		Address:     u.Address,
		Github:      u.Github,
		Linkedin:    u.Linkedin,
		PersonalWeb: u.PersonalWeb,
	}
}
