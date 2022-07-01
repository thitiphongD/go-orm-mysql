package main

import (
	"go-orm-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:daew@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Fname: "Obiwan", Lname: "Kenobi", Username: "Jedi"})
	db.Create(&model.User{Fname: "Anakin", Lname: "Skywalker", Username: "Darth Vador"})
	db.Create(&model.User{Fname: "Ookkotsu", Lname: "Yuta", Username: "Jujutsu"})
}
