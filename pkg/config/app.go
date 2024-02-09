package config

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connStr := "root:Password@1@tcp(127.0.0.1:3307)/sid?parseTime=true"

	client, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	//d, err := gorm.Open("mysql", "sidhantpandey:Sidhant@1/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	DB = client

}

func GetDB() *gorm.DB {
	return DB
}
