package params

import "time"

type CreateUser struct {
	Fullname  string    `json:"full_name"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
}

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}
