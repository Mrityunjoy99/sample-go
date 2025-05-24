package jwt

import (
	"fmt"
	"time"

	"github.com/Mrityunjoy99/sample-go/src/domain/entity"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	jwtSecret     string
	expireTimeSec int
}

func (s *jwtService) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(s.expireTimeSec) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *jwtService) ValidateToken(token string) (*entity.JwtToken, error) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return s.validateTokenFromParsedToken(claims)
}

func (s *jwtService) validateTokenFromParsedToken(claims *jwt.Token) (*entity.JwtToken, error) {
	mapClaims, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	userIdClaim, ok := mapClaims["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid userId claim")
	}

	expClaim, ok := mapClaims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid exp claim")
	}

	return &entity.JwtToken{
		UserId:    userIdClaim,
		ExpiredAt: time.Unix(int64(expClaim), 0),
	}, nil
}
