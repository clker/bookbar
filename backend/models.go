package main

import (
    "github.com/jinzhu/gorm"
    //"encoding/json"
    _"github.com/go-sql-driver/mysql"
    "fmt"
    "strconv"
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
    BookId uint
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


type BookChapter struct{
    Book Book  `json:"book"`
    Chapter Chapter `json:"chapter"`
}

type BookChapters struct{
    Book Book `json:"book"`
    Chapters []Chapter `json:"chapters"`
}

func openDb(){
    db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3311)/bookbar?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
		fmt.Println(err)
	}
    db.AutoMigrate(&BaseModel{},&User{},&Book{},&Chapter{},&Topic{},&Reply{})
}

func closeDb(){
    db.Close()
}

func createBasemodel() uint {
    baseModel := BaseModel{}
    db.Create(&baseModel)
    return baseModel.ID
}

func createOrEditBookModel(book Book) Book {
    fmt.Println("bookid=",book.ID)
    if(book.ID > 0){
        db.Save(&book)
    }else{
        id := createBasemodel()
        book.ID = id
        db.Create(&book)
    }
    return book
}

func deleteBookModel(id string) Book{
    var book Book
    var baseModel BaseModel
    db.Where("ID=?",id).First(&book)
    db.Delete(&book)
    db.Where("ID=?",id).First(&baseModel)
    db.Delete(&baseModel)
    return book
}

func createOrEditChapterModel(id string, chapter Chapter) BookChapter {
    var book Book
    var bookChapter BookChapter
    db.Where("ID=?",id).First(&book)
    if(chapter.ID > 0) {
        db.Save(&chapter)
    }else{
        chapter_id := createBasemodel()
        chapter.ID = chapter_id
        bookId,_ := strconv.Atoi(id)
        chapter.BookId = uint(bookId)
        db.Create(&chapter)
    }
    bookChapter.Book = book
    bookChapter.Chapter = chapter
    return bookChapter
}

func getBookModel(id string) BookChapters {
    var book Book
    var chapters []Chapter
    var bookChapters BookChapters
    db.Where("ID=?",id).First(&book)
    db.Where("book_id=?",id).Find(&chapters)
    //fmt.Println(book)
    //fmt.Println(chapters)
    bookChapters.Book = book
    //jstr,_ := json.Marshal(book)
    //fmt.Println("0",string(jstr))
    bookChapters.Chapters = chapters
    //jstr1,_ := json.Marshal(chapters)
    //fmt.Println("1",string(jstr1))
    //bookChapters := BookChapters{book:book,chapters:make([]Chapter,0)}
    //bookChapters.chapters = make([]Chapter,len(chapters))
    //for k,v := range chapters {
    //    bookChapters.chapters[k] = v
    //}
    //jstr2,_ := json.Marshal(bookChapters)
    return bookChapters
}

