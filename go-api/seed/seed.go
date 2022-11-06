package main

import (
	"project/rapi/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fname    string
	Lname    string
	Username string
	Avatar   string
}

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/project_backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Fname: "Jirawas", Lname: "Tull", Username: "jirawas.t", Avatar: "https://www.melivecode.com/users/1.png"})
	db.Create(&model.User{Fname: "cal", Lname: "ivy", Username: "cal.i", Avatar: "https://www.melivecode.com/users/2.png"})
}
