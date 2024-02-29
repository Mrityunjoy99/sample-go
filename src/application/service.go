package application

import (
	"github.com/Mrityunjoy99/sample-go/src/application/user"
	"github.com/Mrityunjoy99/sample-go/src/infrastructure/repository"
)

type Service struct {
	UserService user.Service
}

func NewService(r repository.Repository) *Service {
	userService := user.NewService(r.UserRepo)
	
	return &Service{
		UserService: userService,
	}
}
