package main
import (
        "github.com/gin-gonic/gin"
        "fmt"
)

func get_books(c *gin.Context){
    var books []Book
    fmt.Println(c)
    db.Find(&books)
	c.JSON(200,books)
}

func get_topics(c *gin.Context){
    var topics []Topic
    db.Find(&topics)
	c.JSON(200,topics)
}

func add_book(c *gin.Context){
    var book Book
    c.BindJSON(&book)
    db.Create(&book)
	c.JSON(200,book)
}
