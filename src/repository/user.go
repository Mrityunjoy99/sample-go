package repository

import (
	"errors"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

//go:generate mockery --name=UserRepository --output=./../mocks/repository --outpkg=mock_repository
type UserRepository interface {
	GetUserById(id uuid.UUID) (*entity.User, genericerror.GenericError)
	CreateUser(user entity.User) (*entity.User, genericerror.GenericError)
	UpdateUser(user entity.User) (*entity.User, genericerror.GenericError)
	DeleteUser(id uuid.UUID) genericerror.GenericError
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetUserById(id uuid.UUID) (*entity.User, genericerror.GenericError) {
	var user entity.User

	result := u.db.Where("id = ?", id).First(&user)
	if (result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound)) || result.RowsAffected == 0 {
		return nil, genericerror.NewGenericError(constant.ErrorCodeResourceNotFound, "user not found", nil, nil)
	}

	if result.Error != nil {
		return nil, genericerror.NewInternalErrByErr(result.Error)
	}

	return &user, nil
}

func (u *userRepository) CreateUser(user entity.User) (*entity.User, genericerror.GenericError) {
	err := u.db.Create(&user).Error
	if err != nil {
		return nil, genericerror.NewInternalErrByErr(err)
	}

	return &user, nil
}

func (u *userRepository) UpdateUser(user entity.User) (*entity.User, genericerror.GenericError) {
	result := u.db.Model(&entity.User{}).Where("id = ?", user.Id).Updates(&user)
	if result.Error != nil {
		return nil, genericerror.NewInternalErrByErr(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, "user not found", nil, nil)
	}

	return &user, nil
}

func (u *userRepository) DeleteUser(id uuid.UUID) genericerror.GenericError {
	result := u.db.Model(&entity.User{}).Delete("id = ?", id)
	if result.Error != nil {
		return genericerror.NewInternalErrByErr(result.Error)
	}

	if result.RowsAffected == 0 {
		return genericerror.NewGenericError(constant.ErrorCodeBadRequest, "user does not exists", nil, nil)
	}

	return nil
}
