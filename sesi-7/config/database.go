package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "127.0.0.1", "3306", "test")
)

func ConnectDB() (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: InitLog(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		log.Print(err)
		panic("Failed to connect database")
	}

	sql, _ := db.DB()

	return db, sql
}

func InitLog() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		},
	)

	return newLogger
}
