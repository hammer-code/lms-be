package domain

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Event struct {
	ID               int       `json:"id"`
	Title            string    `json:"id"`
	Description      string    `json:"description"`
	Author           string    `json:"author"`
	ImageEvent       string    `json:"image_event"`
	DateEvent        null.Time `json:"date_event"`
	Type             string    `json:"type"`
	Location         string    `json:"location"`
	Duration         string    `json:"duration"`
	Capacity         int       `json:"capacity"`
	RegistrationLink string    `json:"registration_link"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        null.Time `json:"updated_at"`
}

func (Event) TableName() string {
	return "events"
}
