package entity

import (
	"time"

	"github.com/Mrityunjoy99/sample-go/src/infrastructure/database"
)

type Url struct {
	database.BaseModel
	OriginalURL string
	ShortURL    string
	ExpiredAt   time.Time
}

func(Url)TableName() string{
	return "url"
}