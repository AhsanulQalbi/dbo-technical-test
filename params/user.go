package params

type CreateUser struct {
	Fullname string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}
