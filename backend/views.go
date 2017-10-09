package main
import (
        "github.com/gin-gonic/gin"
        "fmt"
        //"io/ioutil"
)

func get_books(c *gin.Context){
    var books []Book
    db.Find(&books)
	c.JSON(200,books)
}

func get_topics(c *gin.Context){
    var topics []Topic
    db.Find(&topics)
	c.JSON(200,topics)
}

func create_or_edit_book(c *gin.Context){
    var book Book
    c.BindJSON(&book)
    fmt.Println(book)
    book = create_or_edit_book_model(book)
	c.JSON(200,book)
}

func get_book(c *gin.Context){
    var book Book
    id := c.Param("id")
    db.First(&book,id)
    c.JSON(200,book)
}

func delete_book(c *gin.Context){
    var book Book
    id := c.Param("id")
    book = delete_book_model(id)
    c.JSON(200,book)
}

