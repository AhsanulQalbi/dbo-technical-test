package params

type CreateUser struct {
	Fullname  string `json:"full_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
}

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}
