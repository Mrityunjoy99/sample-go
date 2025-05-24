package jwt

import (
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

	return &entity.JwtToken{
		UserId:    claims.Claims.(jwt.MapClaims)["userId"].(string),
		ExpiredAt: time.Unix(int64(claims.Claims.(jwt.MapClaims)["exp"].(float64)), 0),
	}, nil
}
