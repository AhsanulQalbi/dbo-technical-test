package repositories

import (
	"dbo-technical-test/models"
	"dbo-technical-test/params"
	"fmt"

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
	if queries.Page == 0 {
		queries.Page = 1
	}

	if queries.Size == 0 {
		queries.Size = 10
	}

	fmt.Println(queries.Search)

	var (
		customers []models.Customer
		count     int64
	)
	query := customerRepo.db.
		Order(FULLNAME_ASC).
		Scopes(customerRepo.repoHelpers.Paginate(queries.Page, queries.Size))

	if queries.Search != "" {
		query = query.Where("lower(fullname) ILIKE ?", "%"+queries.Search+"%")
	}
	err := query.Find(&customers).Error

	if err != nil {
		return customers, count, err
	}

	query = customerRepo.db.
		Model(&customers)

	if queries.Search != "" {
		query = query.Where("lower(fullname) ILIKE ?", "%"+queries.Search+"%")
	}
	err = query.Count(&count).
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

func (customerRepo *CustomerRepo) CheckEmail(email string) (*[]models.Customer, error) {
	var customer *[]models.Customer
	err := customerRepo.db.Where("email = ?", email).Find(&customer).Error
	return customer, err
}
