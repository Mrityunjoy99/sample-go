package jwt

import "github.com/Mrityunjoy99/sample-go/src/domain/entity"

//go:generate mockery --name=JwtService --output=./../../../mocks/domain/service --outpkg=mock_domain_service
type JwtService interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*entity.JwtToken, error)
}

func NewJwtService(jwtSecret string, expireTimeSec int) JwtService {
	return &jwtService{jwtSecret: jwtSecret, expireTimeSec: expireTimeSec}
}
