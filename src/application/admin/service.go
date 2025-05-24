package admin

import (
	"time"

	"github.com/Mrityunjoy99/sample-go/src/domain/service/jwt"
)

type Service interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*ValidateTokenRespDto, error)
}

type service struct {
	jwtService jwt.JwtService
}

func NewService(jwtService jwt.JwtService) Service {
	return &service{jwtService: jwtService}
}

func (s *service) GenerateToken(userId string) (string, error) {
	return s.jwtService.GenerateToken(userId)
}

func (s *service) ValidateToken(token string) (*ValidateTokenRespDto, error) {
	jwtToken, err := s.jwtService.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	return &ValidateTokenRespDto{
		UserId:    jwtToken.UserId,
		ExpiredAt: jwtToken.ExpiredAt.Format(time.RFC850),
	}, nil
}
