package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mysql_main() {
	dsn := "root:Long5230413@tcp(127.0.0.1:3306)/archi_faculty?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database ", err)
	}
	db.Name()

	log.Println("Connected to the MySQL database!")
}
