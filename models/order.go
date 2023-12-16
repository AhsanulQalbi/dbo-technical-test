package models

import (
	"time"
)

type Order struct {
	Id              int       `gorm:"not null;uniqueIndex;primaryKey;" json:"id"`
	OrderName       string    `gorm:"not null;" json:"order_name"`
	Quantity        int       `gorm:"not null;" json:"quantity"`
	TotalAmount     int       `gorm:"not null;" json:"total_amount"`
	OrderDate       time.Time `gorm:"not null;" json:"order_date"`
	ShippedDate     time.Time `gorm:"not null;" json:"shipped_date"`
	ArrivedDate     time.Time `gorm:"not null;" json:"arrived_date"`
	ShipAddress     string    `gorm:"not null;" json:"ship_address"`
	ShipProvince    string    `gorm:"not null;" json:"ship_province"`
	ShipCity        string    `gorm:"not null;" json:"ship_city"`
	ShipDistrict    string    `gorm:"not null;" json:"ship_district"`
	ShipSubDistrict string    `gorm:"not null;" json:"ship_sub_district"`
	OrderStatus     string    `gorm:"not null;" json:"order_status"`

	Customer   Customer  `json:"customer"`
	CustomerID string    `gorm:"not null;" json:"customer_id"`
	Product    Product   `json:"product"`
	ProductID  string    `gorm:"not null;" json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
