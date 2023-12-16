package params

type Query struct {
	Status    string `form:"status"`
	Sort      string `form:"sort"`
	SortBy    string `form:"sort_by"`
	OrderBy   string `form:"order_by"`
	Sorting   string
	Search    string `form:"search"`
	SearchBy  string `form:"search_by"`
	Page      int    `form:"page"`
	Size      int    `form:"size"`
	Type      string `form:"type"`
	Date      string `form:"date"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
