package repositories

import (
	"dbo-technical-test/models"
	"dbo-technical-test/params"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db          *gorm.DB
	repoHelpers RepoHelpers
}

func NewOrderRepo(db *gorm.DB, repoHelpers RepoHelpers) *OrderRepo {
	return &OrderRepo{db, repoHelpers}
}

func (orderRepo *OrderRepo) CreateOrder(order models.Order) (*models.Order, error) {
	err := orderRepo.db.Create(&order).Error
	return &order, err
}

func (orderRepo *OrderRepo) GetOrderList(queries *params.Query) ([]models.Order, int64, error) {
	if queries.Page == 0 {
		queries.Page = 1
	}

	if queries.Size == 0 {
		queries.Size = 10
	}
	var (
		orders []models.Order
		count  int64
	)

	query := orderRepo.db.
		Order(ORDER_NAME_ASC).
		Scopes(orderRepo.repoHelpers.Paginate(queries.Page, queries.Size)).
		Preload("Customer").
		Preload("Product")

	if queries.Search != "" {
		query = query.Where("lower(order_name) ILIKE ?", "%"+queries.Search+"%")
	}
	err := query.Find(&orders).Error

	if err != nil {
		return orders, count, err
	}

	query = orderRepo.db.
		Model(&orders)

	if queries.Search != "" {
		query = query.Where("lower(order_name) ILIKE ?", "%"+queries.Search+"%")
	}
	err = query.Count(&count).
		Error

	return orders, count, err
}

func (orderRepo *OrderRepo) GetOrderById(id int) (*models.Order, error) {
	var order models.Order
	err := orderRepo.db.Preload("Customer").
		Preload("Product").Where(WHERE_ID, id).First(&order).Error
	return &order, err
}

func (orderRepo *OrderRepo) UpdateOrder(orderId int, order models.Order) (*models.Order, error) {
	var orderRes models.Order
	err := orderRepo.db.Where(WHERE_ID, orderId).Updates(&order).
		Find(&orderRes).Error
	return &orderRes, err
}

func (orderRepo *OrderRepo) DeleteOrder(orderId int) error {
	var order models.Order
	err := orderRepo.db.Where(WHERE_ID, orderId).Delete(&order).Error
	return err
}
