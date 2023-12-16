package repositories

import (
	"dbo-technical-test/models"
	"dbo-technical-test/params"

	"gorm.io/gorm"
)

type UserRepo struct {
	db          *gorm.DB
	repoHelpers RepoHelpers
}

func NewUserRepo(db *gorm.DB, repoHelpers RepoHelpers) *UserRepo {
	return &UserRepo{db, repoHelpers}
}

func (userRepo *UserRepo) CreateUser(user models.User) (*models.User, error) {
	err := userRepo.db.Create(&user).Error
	return &user, err
}

func (userRepo *UserRepo) GetUserList(queries *params.Query) ([]models.User, int64, error) {
	var (
		users []models.User
		count int64
	)
	err := userRepo.db.
		Order(NAME_ASC).
		Scopes(userRepo.repoHelpers.Paginate(queries.Page, queries.Size)).
		Where("lower("+queries.SearchBy+") ILIKE ?", "%"+queries.Search+"%").
		Find(&users).
		Error

	if err != nil {
		return users, count, err
	}

	err = userRepo.db.
		Model(&users).
		Where("lower("+queries.SearchBy+") ILIKE ?", "%"+queries.Search+"%").
		Count(&count).
		Error

	return users, count, err
}

func (userRepo *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := userRepo.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
