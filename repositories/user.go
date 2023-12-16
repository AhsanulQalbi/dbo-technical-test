package repositories

import (
	"dbo-technical-test/models"

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

func (userRepo *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := userRepo.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (userRepo *UserRepo) ListByEmail(email string) (*[]models.User, error) {
	var user []models.User
	err := userRepo.db.Where("email = ?", email).Find(&user).Error
	return &user, err
}
