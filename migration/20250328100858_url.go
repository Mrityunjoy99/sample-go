package migration

import (
	"time"

	"github.com/Mrityunjoy99/sample-go/src/infrastructure/database"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type V20250328100858Url struct {
	database.BaseModel
	OriginalURL string
	ShortURL    string
	ExpiredAt   time.Time
}

func(V20250328100858Url)TableName() string{
	return "url"
}

var V20250328100858 *gormigrate.Migration = &gormigrate.Migration{
	ID: "20250328100858_url",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&V20250328100858Url{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("url")
	},
}
