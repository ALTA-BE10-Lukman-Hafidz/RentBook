package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Setting struct {
	DBUser string
	DBPass string
}

func ReadEnv() Setting {
	res := Setting{}

	godotenv.Load("local.env")
	res.DBPass = os.Getenv("DBPassword")
	res.DBUser = os.Getenv("DBUser")

	return res
}

func InitDB(s Setting) *gorm.DB {
	conString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/RentBook?charset=utf8mb4&parseTime=True", s.DBUser, s.DBPass)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}
