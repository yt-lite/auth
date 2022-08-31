package db

import (
	"fmt"

	"github.com/yt-lite/auth/config"
	"github.com/yt-lite/libs/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Cfg.PGHost, config.Cfg.PGUser, config.Cfg.PGPassword, config.Cfg.PGDB, config.Cfg.PGPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Fatalf("Error opening database: %s", err)
	}
	logger.Infoln("Postgre connected")
	db.AutoMigrate()
	logger.Infoln("Postgre migrated")
	DB = db
}


