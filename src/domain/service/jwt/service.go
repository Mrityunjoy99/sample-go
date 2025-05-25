package jwt

import (
	"errors"
	"time"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/Mrityunjoy99/sample-go/src/tools/genericerror"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	jwtSecret string
}

func (s *jwtService) GenerateToken(jwtToken *entity.JwtToken) (string, genericerror.GenericError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   jwtToken.UserId,
		"userType": jwtToken.UserType,
		"exp":      jwtToken.ExpiredAt.Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", genericerror.NewInternalErrByErr(err)
	}

	return tokenString, nil
}

func (s *jwtService) ValidateToken(token string) (*entity.JwtToken, genericerror.GenericError) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, err.Error(), nil, err)
		}

		return nil, genericerror.NewInternalErrByErr(err)
	}

	return s.validateTokenFromParsedToken(claims)
}

func (s *jwtService) validateTokenFromParsedToken(claims *jwt.Token) (*entity.JwtToken, genericerror.GenericError) {
	mapClaims, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return nil, genericerror.NewInternalErrByErr(errors.New("invalid token claims"))
	}

	userIdClaim, ok := mapClaims["userId"].(string)
	if !ok {
		return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, "invalid userId claim", nil, errors.New("invalid userId claim"))
	}

	userTypeClaim, ok := mapClaims["userType"].(float64)
	if !ok {
		return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, "invalid userType claim", nil, errors.New("invalid userId claim"))
	}

	expClaim, ok := mapClaims["exp"].(float64)
	if !ok {
		return nil, genericerror.NewGenericError(constant.ErrorCodeBadRequest, "invalid exp claim", nil, errors.New("invalid exp claim"))
	}

	return &entity.JwtToken{
		UserId:    userIdClaim,
		UserType:  constant.UserType(userTypeClaim),
		ExpiredAt: time.Unix(int64(expClaim), 0),
	}, nil
}
