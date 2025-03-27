package user

import (
	"github.com/Mrityunjoy99/sample-go/src/repository"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
	"github.com/google/uuid"
)

type service struct {
	userRepo repository.UserRepository
}

type Service interface {
	GetUserById(id uuid.UUID) (UserResponseDto, genericerror.GenericError)
	CreateUser(dto CreateUserDto) (UserResponseDto, genericerror.GenericError)
	UpdateUser(id uuid.UUID, dto UpdateUserDto) (UserResponseDto, genericerror.GenericError)
	DeleteUser(id uuid.UUID) genericerror.GenericError
}

func NewService(userRepo repository.UserRepository) Service {
	return &service{userRepo: userRepo}
}

func (s *service) GetUserById(id uuid.UUID) (UserResponseDto, genericerror.GenericError) {
	user, gerr := s.userRepo.GetUserById(id)
	if gerr != nil {
		return UserResponseDto{}, gerr
	}

	return newDtoFromEntity(*user), nil
}

func (s *service) CreateUser(dto CreateUserDto) (UserResponseDto, genericerror.GenericError) {
	user := dto.toUserEntity()

	createdUser, gerr := s.userRepo.CreateUser(user)
	if gerr != nil {
		return UserResponseDto{}, gerr
	}

	return newDtoFromEntity(*createdUser), nil
}

func (s *service) UpdateUser(id uuid.UUID, dto UpdateUserDto) (UserResponseDto, genericerror.GenericError) {
	user := dto.toUserEntity()
	user.Id = id

	updatedUser, gerr := s.userRepo.UpdateUser(user)
	if gerr != nil {
		return UserResponseDto{}, gerr
	}

	return newDtoFromEntity(*updatedUser), nil
}

func (s *service) DeleteUser(id uuid.UUID) genericerror.GenericError {
	return s.userRepo.DeleteUser(id)
}
