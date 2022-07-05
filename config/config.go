package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	conString := "root:diangolang15@tcp(localhost:3306)/RentBook?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}
