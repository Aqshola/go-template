package config_db

import (
	"fmt"
	config_general "go-template/src/config/general"
	constant_db "go-template/src/constant/db"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	MySQL *gorm.DB
}

func New(conf config_general.AllConfig) (*Connection, error) {
	dbConfig := conf.DBConfig
	mySqlConfig := dbConfig[constant_db.DB_CONNECTION_MYSQL]

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		mySqlConfig.Username,
		mySqlConfig.Password,
		mySqlConfig.Host,
		mySqlConfig.Port,
		mySqlConfig.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Optional: Get the generic SQL DB to use `Ping`
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return &Connection{
		MySQL: db,
	}, nil

}
