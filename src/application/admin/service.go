package admin

import (
	"time"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/service/jwt"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
)

type Service interface {
	GenerateToken(userId string, userType string) (string, genericerror.GenericError)
	ValidateToken(token string) (*ValidateTokenRespDto, genericerror.GenericError)
}

type service struct {
	jwtService jwt.JwtService
}

func NewService(jwtService jwt.JwtService) Service {
	return &service{jwtService: jwtService}
}

func (s *service) GenerateToken(userId string, userType string) (string, genericerror.GenericError) {
	userTypeEnum, err := constant.GetUserType(userType)
	if err != nil {
		return "", genericerror.NewGenericError(constant.ErrorCodeBadRequest, err.Error(), nil, err)
	}

	return s.jwtService.GenerateToken(userId, userTypeEnum)
}

func (s *service) ValidateToken(token string) (*ValidateTokenRespDto, genericerror.GenericError) {
	jwtToken, err := s.jwtService.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	return &ValidateTokenRespDto{
		UserId:    jwtToken.UserId,
		ExpiredAt: jwtToken.ExpiredAt.Format(time.RFC850),
	}, nil
}
