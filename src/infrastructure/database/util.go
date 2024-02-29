package database

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDialect string

const (
	DialectPostgres GormDialect = "DialectPostgres"
)

type DBConfig struct {
	AppName  string `validate:"required"`
	Host     string `validate:"required"`
	Port     int    `validate:"required"`
	DBName   string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`

	MaxIdleConnCount int `validate:"gte=0"`
	MaxOpenConnCount int `validate:"gtefield=MaxIdleConnCount"`
	// ConnMaxIdleTimeSec sets the maximum amount of time in seconds a connection may be reused.
	ConnMaxIdleTimeSec int `validate:"gte=0"`

	// ConnMaxLifeTimeSec sets the maximum amount of time in seconds a connection may be reused
	ConnMaxLifeTimeSec int `validate:"gtefield=ConnMaxIdleTimeSec"`

	Dialect GormDialect `validate:"required"`

	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction *bool `default:"true"`

	Options *gorm.Config
}

func ConnectDatabase(c DBConfig) (*gorm.DB, error) {
	const (
		RepositoryPostgresConnectWaitSec = 3
		ConnectRetryAttempts             = 10
	)

	for i := 0; i < ConnectRetryAttempts; i++ {
		db, err := dialDb(c)
		if err != nil {
			time.Sleep(RepositoryPostgresConnectWaitSec * time.Second)
			continue
		} else {
			return db, nil
		}
	}

	return nil, errors.New("unable to connect to database")
}

func dialDb(c DBConfig) (*gorm.DB, error) {
	dialector, serr := getDialector(c)
	if serr != nil {
		return nil, serr
	}

	dbObj, err := gorm.Open(
		dialector,
		c.Options,
	)

	if err != nil {
		return nil, errors.New("unable to connect to database")
	}

	if db, err := dbObj.DB(); err == nil {
		// pooling setup
		if c.MaxIdleConnCount > 0 {
			db.SetMaxIdleConns(c.MaxIdleConnCount)
		}

		if c.MaxOpenConnCount > 0 {
			db.SetMaxOpenConns(c.MaxOpenConnCount)
		}

		// conn timeout setup
		if c.ConnMaxIdleTimeSec > 0 {
			db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTimeSec) * time.Second)
		}

		if c.ConnMaxLifeTimeSec > 0 {
			db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTimeSec) * time.Second)
		}
	} else {
		return nil, err
	}

	return dbObj, nil
}

func getDialector(c DBConfig) (gorm.Dialector, error) {
	switch c.Dialect {
	case DialectPostgres:
		return getPGDialector(c), nil

	default:
		var dialector gorm.Dialector
		return dialector, errors.New("invalid dialect")
	}
}

func getPGDialector(c DBConfig) gorm.Dialector {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s application_name=%s",
		c.Host,
		c.Port,
		c.User,
		c.DBName,
		c.Password,
		c.AppName,
	)

	return postgres.Open(dsn)
}
