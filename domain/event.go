package domain

import (
	"context"
	"net/http"
	"time"

	"gopkg.in/guregu/null.v4"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, data Event) (uint, error)
	// UpdateEvent(ctx context.Context, payload Event) error
	// DeleteEvent(ctx context.Context, id int) error
	// GetEventByID(ctx context.Context, id int) (Event, error)
	// GetEvents(ctx context.Context, filter EventFilter) ([]Event, error)
	CreateEventTag(ctx context.Context, data EventTag) (uint, error)
	CreateEventSpeaker(ctx context.Context, data EventSpeaker) (uint, error)
	GetEvents(ctx context.Context, filter EventFilter) (tData int, data []Event, err error)
	CreatePayEvent(ctx context.Context, event EventPay) (uint, error)
	CreateRegisterEvent(ctx context.Context, event RegistrationEvent) (uint, error)
	GetEvent(ctx context.Context, eventID uint) (data Event, err error)
	GetRegistrationEvent(ctx context.Context, orderNo string) (data RegistrationEvent, err error)
}

type EventUsecase interface {
	CreateEvent(ctx context.Context, payload CreateEvenPayload) error
	GetEvents(ctx context.Context, filter EventFilter) (data []Event, pagination Pagination, err error)
	CreateRegisterEvent(ctx context.Context, payload RegisterEventPayload) (RegisterEventResponse, error)
	CreatePayEvent(ctx context.Context, payload EventPayPayload) error
	GetEventByID(ctx context.Context, id uint) (resp Event, err error)
}

type EventHandler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	GetEvents(w http.ResponseWriter, r *http.Request)
	RegisterEvent(w http.ResponseWriter, r *http.Request)
	PayEvent(w http.ResponseWriter, r *http.Request)
	GetEventByID(w http.ResponseWriter, r *http.Request)
}

type Event struct {
	ID               uint           `json:"id" gorm:"primarykey"`
	Title            string         `json:"title" `
	Description      string         `json:"description"`
	Author           string         `json:"author"`
	ImageEvent       string         `json:"image_event"`
	DateEvent        null.Time      `json:"date_event"`
	Type             string         `json:"type"`
	Location         string         `json:"location"`
	Duration         string         `json:"duration"`
	Capacity         int            `json:"capacity"`
	Status           string         `json:"status"`                                          // comming soon
	Tags             []EventTag     `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;"` // Ensure foreign key is correctly referenced
	Speakers         []EventSpeaker `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;"` // Ensure foreign key is correctly referenced
	RegistrationLink string         `json:"registration_link"`
	Price            float64        `json:"price"` // 0 == free
	CreatedByUserID  int            `json:"created_by_user_id"`
	UpdatedByUserID  int            `json:"updated_by_user_id"`
	DeletedByUserID  int            `json:"deleted_by_user_id"`
	BookingStart     null.Time      `json:"booking_start"`
	BookingEnd       null.Time      `json:"booking_end"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        null.Time      `json:"updated_at"`
	DeletedAt        null.Time      `json:"deleted_at"`
}

func (Event) TableName() string {
	return "events"
}

type EventTag struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	EventID uint   `json:"event_id"`
	Tag     string `json:"tags"`
}

func (EventTag) TableName() string {
	return "event_tags"
}

type EventSpeaker struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	EventID uint   `json:"event_id"`
	Name    string `json:"name"`
}

func (EventSpeaker) TableName() string {
	return "event_speakers"
}

type CreateEvenPayload struct {
	Title            string    `json:"title" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	Author           string    `json:"author" validate:"required"`
	FileName         string    `json:"file_name" validate:"required"`
	DateEvent        null.Time `json:"date_event" validate:"required"`
	Type             string    `json:"type" validate:"required"`
	Location         string    `json:"location" validate:"required"`
	Duration         string    `json:"duration" validate:"required"`
	Status           string    `json:"status" validate:"required"`
	Capacity         int       `json:"capacity" validate:"required"`
	Price            float64   `json:"price"`
	RegistrationLink string    `json:"registration_link"`
	Tags             []string  `json:"tags"`
	Speakers         []string  `json:"speakers"`
	BookingStart     null.Time `json:"booking_start"`
	BookingEnd       null.Time `json:"booking_end"`
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
	ID        uint
	Title     string
	Type      string
	Status    string
	StartDate null.Time
	EndDate   null.Time
	FilterPagination
}

type EventPay struct {
	ID                uint    `json:"id" gorm:"primarykey"`
	RegistrationEvent uint    `json:"registration_event_id"`
	ImageProofPayment string  `json:"image_proof_payment"`
	NetAmount         float64 `json:"net_amount"`
}

func (EventPay) TableName() string {
	return "event_pays"
}

type RegisterEventPayload struct {
	EventID     uint   `json:"event_id"`
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type RegistrationEvent struct {
	ID                uint      `json:"id" gorm:"primarykey"`
	OrderNo           string    `json:"order_no"`
	EventID           uint      `json:"event_id"` // lock event
	UserID            string    `json:"user_id"`  // lock user
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phone_number"`
	ImageProofPayment string    `json:"image_proof_payment"`
	PaymentDate       null.Time `json:"payment_date"`
	Status            string    `json:"status"` // register, pay, approve/cancel/decline
	UpToYou           string    `json:"up_to_you"`
	CreatedByUserID   int       `json:"created_by_user_id"`
	UpdatedByUserID   int       `json:"updated_by_user_id"`
	DeletedByUserID   int       `json:"deleted_by_user_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         null.Time `json:"updated_at"`
	DeletedAt         null.Time `json:"deleted_at"`
}

func (RegistrationEvent) TableName() string {
	return "registration_events"
}

type RegistrationEventDTO struct {
	OrderNo string `json:"order_no"`
}

type RegisterEventResponse struct {
	OrderNo string `json:"order_no"`
}

type EventPayPayload struct {
	OrderNo           string  `json:"order_no"`
	ImageProofPayment string  `json:"image_proof_payment"`
	NetAmount         float64 `json:"net_amount"`
}
