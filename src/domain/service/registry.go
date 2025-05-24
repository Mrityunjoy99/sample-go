package service

import (
	"github.com/Mrityunjoy99/sample-go/src/common/config"
	"github.com/Mrityunjoy99/sample-go/src/domain/service/jwt"
)

type ServiceRegistry struct {
	JwtService jwt.JwtService
}

func NewServiceRegistry(config *config.Config) *ServiceRegistry {
	if config == nil {
		panic("config is required")
	}

	if config.Jwt.Secret == "" {
		panic("JWT secret is required")
	}

	if config.Jwt.ExpireTimeSec <= 0 {
		panic("JWT expiration time must be positive")
	}
	
	jwtService := jwt.NewJwtService(config.Jwt.Secret, config.Jwt.ExpireTimeSec)
	return &ServiceRegistry{JwtService: jwtService}
}
