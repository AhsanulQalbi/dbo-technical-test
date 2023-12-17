package params

import "time"

type CreateCustomer struct {
	Email     string `json:"email"`
	Fullname  string `json:"full_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
}

type UpdateCustomer struct {
	Fullname  string `json:"full_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
}

type DetailCustomer struct {
	Id        int       `json:"id"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	BirthDate string    `json:"birth_date"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerResponse struct {
	Message string         `json:"message"`
	Data    DetailCustomer `json:"data"`
}
