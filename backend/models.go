package main

import (
    "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
    "fmt"
)

var db *gorm.DB
var err error

type User struct {
    gorm.Model
    Username string `gorm:"not null" form:"username" json:"username"`
}

type Book struct {
    gorm.Model
    Name string `gorm:"not null" form:"name" json:"name"`
    Description string `gorm:"" form:"description" json:"description"`
}

type Topic struct {
    gorm.Model
    Title string `gorm:"not null" form:"title" json:"title"`
    Content string `gorm:"" form:"content" json:"content"`
}

func open_db(){
    db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3311)/bookbar?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
		fmt.Println(err)
	}
    db.AutoMigrate(&User{},&Book{},)
}

func close_db(){
    db.Close()
}
