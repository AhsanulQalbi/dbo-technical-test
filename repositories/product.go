package repositories

import (
	"dbo-technical-test/models"
	"dbo-technical-test/params"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db          *gorm.DB
	repoHelpers RepoHelpers
}

func NewProductRepo(db *gorm.DB, repoHelpers RepoHelpers) *ProductRepo {
	return &ProductRepo{db, repoHelpers}
}

func (productRepo *ProductRepo) CreateProduct(product models.Product) (*models.Product, error) {
	err := productRepo.db.Create(&product).Error
	return &product, err
}

func (productRepo *ProductRepo) GetProductList(queries *params.Query) ([]params.ListProduct, int64, error) {
	if queries.Page == 0 {
		queries.Page = 1
	}

	if queries.Size == 0 {
		queries.Size = 10
	}

	var (
		products []params.ListProduct
		count    int64
	)

	query := productRepo.db.
		Order(NAME_ASC).
		Scopes(productRepo.repoHelpers.Paginate(queries.Page, queries.Size))

	if queries.Search != "" {
		query = query.Where("lower(name) ILIKE ?", "%"+queries.Search+"%")
	}
	err := query.Model(models.Product{}).Find(&products).Error

	if err != nil {
		return products, count, err
	}

	query = productRepo.db.
		Model(models.Product{})

	if queries.Search != "" {
		query = query.Where("lower(name) ILIKE ?", "%"+queries.Search+"%")
	}
	err = query.Count(&count).
		Error
	return products, count, err
}

func (productRepo *ProductRepo) GetProductById(id int) (*models.Product, error) {
	var product models.Product
	err := productRepo.db.Where(WHERE_ID, id).First(&product).Error
	return &product, err
}

func (productRepo *ProductRepo) UpdateProduct(productId int, product models.Product) (*models.Product, error) {
	var productRes models.Product
	err := productRepo.db.Where(WHERE_ID, productId).Updates(&product).
		Find(&productRes).Error
	return &productRes, err
}

func (productRepo *ProductRepo) DeleteProduct(productId int) error {
	var product models.Product
	err := productRepo.db.Where(WHERE_ID, productId).Delete(&product).Error
	return err
}
