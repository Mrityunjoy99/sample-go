package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepo UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepo := NewUserRepository(db)

	return &Repository{
		UserRepo: userRepo,
	}
}
