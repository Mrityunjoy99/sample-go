package application

import (
	"github.com/Mrityunjoy99/sample-go/src/application/url"
	"github.com/Mrityunjoy99/sample-go/src/application/user"
	"github.com/Mrityunjoy99/sample-go/src/repository"
)

type Service struct {
	UserService user.Service
	UrlService  url.Service
}

func NewService(r repository.Repository) *Service {
	userService := user.NewService(r.UserRepo)
	urlService := url.NewService(r.UrlRepo)

	return &Service{
		UserService: userService,
		UrlService:  urlService,
	}
}
