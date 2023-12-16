package models

import "time"

type Product struct {
	Id          int    `gorm:"not null;uniqueIndex;primaryKey;" json:"id"`
	Name        string `gorm:"not null;" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	Price       int    `gorm:"not null;" json:"price"`
	Weight      int    `json:"weight"`
	ImageUrl    string `json:"image_url"`
	Stock       int    `gorm:"not null;default:0" json:"stock"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
