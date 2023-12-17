package repositories

import "gorm.io/gorm"

type RepoHelpers struct{}

func NewRepoHelpers() *RepoHelpers {
	return &RepoHelpers{}
}

func (repoHelpers *RepoHelpers) Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
