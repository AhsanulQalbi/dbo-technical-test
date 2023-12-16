package models

import "time"

type Customer struct {
	Id        int       `gorm:"not null;uniqueIndex;primaryKey;" json:"id"`
	Fullname  string    `gorm:"not null;" json:"full_name"`
	Password  string    `gorm:"not null;" json:"-"`
	Email     string    `gorm:"not null;" json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	LastLogin time.Time `json:"last_login"`
	Status    string    `gorm:"not null;default:'Active'" json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
