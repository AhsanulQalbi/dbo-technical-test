package repositories

import (
	"dbo-technical-test/models"
	"dbo-technical-test/params"

	"gorm.io/gorm"
)

type CustomerRepo struct {
	db          *gorm.DB
	repoHelpers RepoHelpers
}

func NewCustomerRepo(db *gorm.DB, repoHelpers RepoHelpers) *CustomerRepo {
	return &CustomerRepo{db, repoHelpers}
}

func (customerRepo *CustomerRepo) CreateCustomer(customer models.Customer) (*models.Customer, error) {
	err := customerRepo.db.Create(&customer).Error
	return &customer, err
}

func (customerRepo *CustomerRepo) GetCustomerList(queries *params.Query) ([]models.Customer, int64, error) {
	var (
		customers []models.Customer
		count     int64
	)
	err := customerRepo.db.
		Order(NAME_ASC).
		Scopes(customerRepo.repoHelpers.Paginate(queries.Page, queries.Size)).
		Where("lower(fullname) ILIKE ?", "%"+queries.Search+"%").
		Find(&customers).
		Error

	if err != nil {
		return customers, count, err
	}

	err = customerRepo.db.
		Model(&customers).
		Where("lower(fullname) ILIKE ?", "%"+queries.Search+"%").
		Count(&count).
		Error

	return customers, count, err
}

func (customerRepo *CustomerRepo) GetCustomerById(id int) (*models.Customer, error) {
	var customer models.Customer
	err := customerRepo.db.Where(WHERE_ID, id).First(&customer).Error
	return &customer, err
}

func (customerRepo *CustomerRepo) UpdateCustomer(customerId int, customer models.Customer) (*models.Customer, error) {
	var customerRes models.Customer
	err := customerRepo.db.Where(WHERE_ID, customerId).Updates(&customer).
		Find(&customerRes).Error
	return &customerRes, err
}

func (customerRepo *CustomerRepo) DeleteCustomer(customerId int) error {
	var customer models.Customer
	err := customerRepo.db.Where(WHERE_ID, customerId).Delete(&customer).Error
	return err
}
