package service

import (
	"errors"

	"github.com/Mrityunjoy99/sample-go/src/common/config"
	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/service/jwt"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
)

type ServiceRegistry struct {
	JwtService jwt.JwtService
}

func NewServiceRegistry(config *config.Config) (*ServiceRegistry, genericerror.GenericError) {
	if config == nil {
		return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, "config is required", nil, errors.New("config is required"))
	}

	if config.Jwt.Secret == "" {
		return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, "JWT secret is required", nil, errors.New("JWT secret is required"))
	}

	if config.Jwt.ExpireTimeSec <= 0 {
		return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, "JWT expiration time must be positive", nil, errors.New("JWT expiration time must be positive"))
	}

	jwtService := jwt.NewJwtService(config.Jwt.Secret)
	return &ServiceRegistry{JwtService: jwtService}, nil
}
