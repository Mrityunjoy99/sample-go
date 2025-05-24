package service

import (
	"github.com/Mrityunjoy99/sample-go/src/common/config"
	"github.com/Mrityunjoy99/sample-go/src/domain/service/jwt"
)

type ServiceRegistry struct {
	JwtService jwt.JwtService
}

func NewServiceRegistry(config *config.Config) *ServiceRegistry {
	jwtService := jwt.NewJwtService(config.Jwt.Secret, config.Jwt.ExpireTimeSec)
	return &ServiceRegistry{JwtService: jwtService}
}
