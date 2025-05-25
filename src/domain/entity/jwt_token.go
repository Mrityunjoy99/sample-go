package entity

import (
	"time"

	"github.com/Mrityunjoy99/sample-go/src/common/constant"
)

type JwtToken struct {
	UserId    string
	UserType  constant.UserType
	ExpiredAt time.Time
}
