package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id == uuid.Nil {
		uuid := uuid.New()
		base.Id = uuid
	}

	return nil
}

func (base *BaseModel) SetCreatedAt(createdAt time.Time) {
	base.CreatedAt = createdAt
}
