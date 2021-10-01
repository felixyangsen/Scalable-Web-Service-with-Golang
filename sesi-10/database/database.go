package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"myapp/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))

	var postgres = postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	db, err = gorm.Open(postgres, &gorm.Config{
		Logger: InitLog(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		log.Print(err)
		panic("Failed to connect database")
	}

	fmt.Println("connection success")
	db.Debug().AutoMigrate(model.User{}, model.Product{})
}

func GetDB() *gorm.DB {
	return db
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
