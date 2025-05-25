package admin

import (
	"time"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/Mrityunjoy99/sample-go/src/domain/service/jwt"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
)

type Service interface {
	GenerateToken(userId string, userType string) (string, genericerror.GenericError)
}

type service struct {
	jwtService    jwt.JwtService
	expireTimeSec int
}

func NewService(jwtService jwt.JwtService, expireTimeSec int) Service {
	return &service{jwtService: jwtService, expireTimeSec: expireTimeSec}
}

func (s *service) GenerateToken(userId string, userType string) (string, genericerror.GenericError) {
	userTypeEnum, err := constant.GetUserType(userType)
	if err != nil {
		return "", genericerror.NewGenericError(constant.ErrorCodeBadRequest, err.Error(), nil, err)
	}

	jwtTokenEntity := &entity.JwtToken{
		UserId:    userId,
		UserType:  userTypeEnum,
		ExpiredAt: time.Now().Add(time.Duration(s.expireTimeSec) * time.Second),
	}

	return s.jwtService.GenerateToken(jwtTokenEntity)
}
