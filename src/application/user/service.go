package user

import (
	"github.com/Mrityunjoy99/sample-go/src/infrastructure/repository"
	"github.com/google/uuid"
)

type service struct {
	userRepo repository.UserRepository
}

type Service interface {
	GetUserById(id uuid.UUID) (UserResponseDto, error)
	CreateUser(dto CreateUserDto) (UserResponseDto, error)
	UpdateUser(dto UpdateUserDto) (UserResponseDto, error)
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

func (s *service) UpdateUser(dto UpdateUserDto) (UserResponseDto, error) {
	currentUser, err := s.userRepo.GetUserById(dto.Id)
	if err != nil {
		return UserResponseDto{}, err
	}

	currentUser.FirstName = dto.FirstName
	currentUser.LastName = dto.LastName
	currentUser.Email = dto.Email
	currentUser.Phone = dto.Phone

	updatedUser, err := s.userRepo.UpdateUser(currentUser)
	if err != nil {
		return UserResponseDto{}, err
	}
	return newDtoFromEntity(updatedUser), nil
}
