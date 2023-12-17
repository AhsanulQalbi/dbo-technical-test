package models

import "time"

type User struct {
	Id        int       `gorm:"not null;uniqueIndex;primaryKey;" json:"id"`
	Fullname  string    `gorm:"not null;" json:"full_name"`
	Password  string    `gorm:"not null;" json:"-"`
	Email     string    `gorm:"unique;not null;" json:"email"`
	Role      string    `gorm:"not null;default:'Admin'" json:"role"`
	LastLogin time.Time `json:"last_login"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
