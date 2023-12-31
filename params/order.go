package params

type CreateOrder struct {
	OrderName       string `json:"order_name"`
	Quantity        int    `json:"quantity"`
	TotalAmount     int    `json:"total_amount"`
	OrderDate       string `json:"order_date"`
	ShippedDate     string `json:"shipped_date"`
	ArrivedDate     string `json:"arrived_date"`
	ShipAddress     string `json:"ship_address"`
	ShipProvince    string `json:"ship_province"`
	ShipCity        string `json:"ship_city"`
	ShipDistrict    string `json:"ship_district"`
	ShipSubDistrict string `json:"ship_sub_district"`
	OrderStatus     string `json:"order_status"`
	CustomerID      int    `json:"customer_id"`
	ProductID       int    `json:"product_id"`
}

type UpdateOrder struct {
	ShippedDate     string `json:"shipped_date"`
	ArrivedDate     string `json:"arrived_date"`
	ShipAddress     string `json:"ship_address"`
	ShipProvince    string `json:"ship_province"`
	ShipCity        string `json:"ship_city"`
	ShipDistrict    string `json:"ship_district"`
	ShipSubDistrict string `json:"ship_sub_district"`
	OrderStatus     string `json:"order_status"`
}

type OrderResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// type DetailOrder struct {
// 	Id              int       `json:"id"`
// 	OrderName       string    `json:"order_name"`
// 	Quantity        int       `json:"quantity"`
// 	TotalAmount     int       `json:"total_amount"`
// 	OrderDate       time.Time `json:"order_date"`
// 	ShippedDate     time.Time `json:"shipped_date"`
// 	ArrivedDate     time.Time `json:"arrived_date"`
// 	ShipAddress     string    `json:"ship_address"`
// 	ShipProvince    string    `json:"ship_province"`
// 	ShipCity        string    `json:"ship_city"`
// 	ShipDistrict    string    `json:"ship_district"`
// 	ShipSubDistrict string    `json:"ship_sub_district"`
// 	OrderStatus     string    `json:"order_status"`
// 	Customer        Customer  `json:"customer"`
// 	CustomerID      int       `gorm:"not null;" json:"customer_id"`
// 	Product         Product   `json:"product"`
// 	ProductID       int       `gorm:"not null;" json:"product_id"`
// 	CreatedAt       time.Time `json:"created_at"`
// 	UpdatedAt       time.Time `json:"updated_at"`
// 	DeletedAt       time.Time `json:"deleted_at"`
// }
