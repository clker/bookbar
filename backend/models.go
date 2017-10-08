package main

import (
    "github.com/jinzhu/gorm"
    _"github.com/go-sql-driver/mysql"
    "fmt"
)

var db *gorm.DB
var err error

type BaseModel struct {
    gorm.Model
}

type User struct {
    gorm.Model
    Username string `gorm:"not null" json:"username"`
    Password string `gorm:"not null" json:"password"`
}

type Book struct {
    gorm.Model
    User User
    UserID uint `gorm:"not null" json:"userid"`
    Name string `gorm:"not null" json:"name"`
    Description string `gorm:"" json:"description"`
    //BaseModel BaseModel `gorm:"polymorphic:Owner;"`
}

type Chapter struct {
    gorm.Model
    Book Book
    BookID uint
    Title string `gorm:"not null" json:"title"`
    Content string `gorm:"" json:"content"`
    BaseModel BaseModel `gorm:"polymorphic:Owner;"`
    Order uint
}

type Topic struct {
    gorm.Model
    Title string `gorm:"not null" json:"title"`
    Content string `gorm:"" json:"content"`
    BaseModel BaseModel `gorm:"polymorphic:Owner;"`
}

type Reply struct {
    gorm.Model
    Content string `gorm:"" json:"content"`
    BaseModel BaseModel `gorm:"polymorphic:Owner;"`
    Topic Topic
    TopicID uint
}

func open_db(){
    db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3311)/bookbar?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
		fmt.Println(err)
	}
    db.AutoMigrate(&BaseModel{},&User{},&Book{},&Chapter{},&Topic{},&Reply{})
}

func close_db(){
    db.Close()
}

func create_basemodel() uint {
    baseModel := BaseModel{}
    db.Create(&baseModel)
    return baseModel.ID
}

func create_or_edit_book_model(book Book) Book {
    fmt.Println("bookid=",book.ID)
    if(book.ID > 0){
        db.Save(&book)
    }else{
        id := create_basemodel()
        book.ID = id
        db.Create(&book)
    }
    return book
}

func delete_book_model(id string) Book{
    var book Book
    var baseModel BaseModel
    db.Where("ID=?",id).First(&book)
    db.Delete(&book)
    db.Where("ID=?",id).First(&baseModel)
    db.Delete(&baseModel)
    db.Where("ID=?",id).First(&book)
    return book
}
