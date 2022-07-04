package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	connstr := "root:@tcp(localhost:3306)/RentBook?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(connstr), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}
