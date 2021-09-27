package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"myapp/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "127.0.0.1", "3306", "test")
	db  *gorm.DB
	err error
)

func StartDB() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
