package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DbUser    string
	DbPass    string
	DbName    string
	DbHost    string
	DbPort    string
	DbCharset string
	DbDriver  string
}

var DB *gorm.DB

func ConnectDB(config DBConfig) (*gorm.DB, error) {
	var err error
	var DSN string

	switch config.DbDriver {
	case "mysql":
		DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbName, config.DbCharset)
		DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	case "postgres":
		DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.DbHost, config.DbUser, config.DbPass, config.DbName, config.DbPort)
		DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unknown database driver")
	}

	if err != nil {
		return nil, err
	}

	return DB, nil
}
