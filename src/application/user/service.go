package user

import (
	"github.com/Mrityunjoy99/sample-go/src/repository"
	"github.com/google/uuid"
)

type service struct {
	userRepo repository.UserRepository
}

type Service interface {
	GetUserById(id uuid.UUID) (UserResponseDto, error)
	CreateUser(dto CreateUserDto) (UserResponseDto, error)
}

func NewService(userRepo repository.UserRepository) Service {
	return &service{userRepo: userRepo}
}

func (s *service) GetUserById(id uuid.UUID) (UserResponseDto, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		return UserResponseDto{}, err
	}
	return newDtoFromEntity(user), nil
}

func (s *service) CreateUser(dto CreateUserDto) (UserResponseDto, error) {
	user := dto.toUserEntity()
	user, err := s.userRepo.CreateUser(user)
	if err != nil {
		return UserResponseDto{}, err
	}
	return newDtoFromEntity(user), nil
}
