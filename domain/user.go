package domain

type (
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
)

// define gorm table of struct
func (User) TableName() string {
	return "users"
}

