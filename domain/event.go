package domain

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, payload Event) (int, error)
	UpdateEvent(ctx context.Context, payload Event) error
	DeleteEvent(ctx context.Context, id int) error
	GetEventByID(ctx context.Context, id int) (Event, error)
	GetEvents(ctx context.Context, filter EventFilter) ([]Event, error)
}

type EventUsecase interface {
	CreateEvent(ctx context.Context, payload CreateEvenPayload) (EventDTO, error)
	UpdateEvent(ctx context.Context, payload UpdateEvenPayload) error
	DeleteEvent(ctx context.Context, id int) error
	GetEventByID(ctx context.Context, id int) (EventDTO, error)
	GetEvents(ctx context.Context, filter EventFilter) ([]EventDTO, error)
}

type Event struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Author           string    `json:"author"`
	ImageEvent       string    `json:"image_event"`
	DateEvent        null.Time `json:"date_event"`
	Type             string    `json:"type"`
	Location         string    `json:"location"`
	Duration         string    `json:"duration"`
	Capacity         int       `json:"capacity"`
	RegistrationLink string    `json:"registration_link"`
	CreatedByUserID  int       `json:"created_by_user_id"`
	UpdatedByUserID  int       `json:"updated_by_user_id"`
	DeletedByUserID  int       `json:"deleted_by_user_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        null.Time `json:"updated_at"`
	DeletedAt        null.Time `json:"deleted_at"`
}

func (Event) TableName() string {
	return "events"
}

type CreateEvenPayload struct {
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
}

type EventDTO struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Author           string    `json:"author"`
	ImageEvent       string    `json:"image_event"`
	DateEvent        null.Time `json:"date_event"`
	Type             string    `json:"type"`
	Location         string    `json:"location"`
	Duration         string    `json:"duration"`
	Capacity         int       `json:"capacity"`
	RegistrationLink string    `json:"registration_link"`
}

type UpdateEvenPayload struct {
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	Author           string    `json:"author"`
	ImageEvent       string    `json:"image_event"`
	DateEvent        null.Time `json:"date_event"`
	Type             string    `json:"type"`
	Location         string    `json:"location"`
	Duration         string    `json:"duration"`
	Capacity         int       `json:"capacity"`
	RegistrationLink string    `json:"registration_link"`
}

type EventFilter struct {
	Title     string
	Type      string
	Status    string
	StartDate null.Time
	EndDDate  null.Time
}
