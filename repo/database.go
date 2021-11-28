package repo

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func GormDB() *gorm.DB {
	return gormDB
}

func InitDBConnection() {
	dsn := "root:@tcp(127.0.0.1:3306)/inspirit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error in setup DB connection ", err)
	}
	log.Println("DB connection sucess")
	gormDB = db
}
