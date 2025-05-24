package application

import (
	"github.com/Mrityunjoy99/sample-go/src/application/admin"
	"github.com/Mrityunjoy99/sample-go/src/application/user"
	"github.com/Mrityunjoy99/sample-go/src/domain/service"
	"github.com/Mrityunjoy99/sample-go/src/repository"
)

type Service struct {
	UserService  user.Service
	AdminService admin.Service
}

func NewService(r *repository.Repository, domainService *service.ServiceRegistry) *Service {
	if domainService == nil {
		panic("domainService is required")
	}

	if domainService.JwtService == nil {
		panic("JwtService is required in domainService")
	}
	
	userService := user.NewService(r.UserRepo)

	return &Service{
		UserService:  userService,
		AdminService: admin.NewService(domainService.JwtService),
	}
}
