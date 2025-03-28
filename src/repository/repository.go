package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepo UserRepository
	UrlRepo  URLRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepo := NewUserRepository(db)
	urlRepo := NewUrlRepository(db)

	return &Repository{
		UserRepo: userRepo,
		UrlRepo:  urlRepo,
	}
}
