package params

type CreateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Weight      int    `json:"weight"`
	ImageUrl    string `json:"image_url"`
	Stock       int    `json:"stock"`
}

type UpdateProduct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Weight      int    `json:"weight"`
	ImageUrl    string `json:"image_url"`
	Stock       int    `json:"stock"`
}

type ListProduct struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageUrl string `json:"image_url"`
	Stock    int    `json:"stock"`
}

type DetailProduct struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Weight      int    `json:"weight"`
	ImageUrl    string `json:"image_url"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductResponse struct {
	Message string        `json:"message"`
	Data    DetailProduct `json:"data"`
}
