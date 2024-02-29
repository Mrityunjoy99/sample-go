package repository

import (
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

//go:generate mockery --name=UserRepository --output=./../../mocks/infrastructure/repository --outpkg=mock_repository
type UserRepository interface {
	GetUserById(id uuid.UUID) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetUserById(id uuid.UUID) (entity.User, error) {
	var user entity.User

	err := u.db.Where("id = ?", id).First(&user).Error

	return user, err
}

func (u *userRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}
