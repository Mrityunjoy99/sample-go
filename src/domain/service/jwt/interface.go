package jwt

import (
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
)

//go:generate mockery --name=JwtService --output=./../../../mocks/domain/service --outpkg=mock_domain_service
type JwtService interface {
	GenerateToken(*entity.JwtToken) (string, genericerror.GenericError)
	ValidateToken(token string) (*entity.JwtToken, genericerror.GenericError)
}

func NewJwtService(jwtSecret string) JwtService {
	return &jwtService{jwtSecret: jwtSecret}
}
