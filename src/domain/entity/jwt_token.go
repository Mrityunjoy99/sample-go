package entity

import "time"

type JwtToken struct {
	UserId    string
	ExpiredAt time.Time
}
