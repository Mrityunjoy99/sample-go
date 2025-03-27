package database

import (
	"errors"
	"fmt"

	"github.com/Mrityunjoy99/sample-go/src/common/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect(c *config.Config) (*gorm.DB, error) {
	logLevel := logger.Info

	gConfig := DBConfig{
		AppName:            c.App.Name,
		Host:               c.DB.Host,
		Port:               c.DB.Port,
		DBName:             c.DB.Name,
		User:               c.DB.User,
		Password:           c.DB.Password,
		Dialect:            DialectPostgres,
		MaxIdleConnCount:   c.DB.MaxIdleConnections,
		MaxOpenConnCount:   c.DB.MaxOpenConnections,
		ConnMaxIdleTimeSec: c.DB.ConnMaxIdleTimeSec,
		ConnMaxLifeTimeSec: c.DB.ConnMaxLifeTimeSec,
		Options: &gorm.Config{
			Logger:                 logger.Default.LogMode(logLevel),
			SkipDefaultTransaction: true,
		},
	}

	fmt.Println("Connecting to database")
	fmt.Println("config: ", gConfig)

	con, err := ConnectDatabase(gConfig)
	if err != nil {
		return nil, err
	}

	db = con

	return db, nil
}

func GetInstance() *gorm.DB {
	return db
}

func SetInstance(g *gorm.DB) {
	db = g
}

func CloseDB() error {
	dbInstance, err := db.DB()
	if err != nil {
		return errors.New("error while getting db.DB() in dbInstance.Close() call")
	}

	closeDBerr := dbInstance.Close()
	if closeDBerr != nil {
		return errors.New("error while closing dbInstance")
	}

	return nil
}
