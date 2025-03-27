package migration

import (
	"github.com/Mrityunjoy99/sample-go/src/infrastructure/database"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type V20240229232650User struct {
	database.BaseModel
	FirstName string `gorm:"column:first_name;not null"`
	LastName  string `gorm:"column:last_name;not null"`
	Email     string `gorm:"column:email;not null"`
	Phone     string `gorm:"column:phone;not null"`
}

func (V20240229232650User) TableName() string {
	return "users"
}

var V20240229232650 *gormigrate.Migration = &gormigrate.Migration{
	ID: "20240229232650_User",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&V20240229232650User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
