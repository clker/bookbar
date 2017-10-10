package main
import (
        "github.com/gin-gonic/gin"
        "fmt"
        //"io/ioutil"
)

func getBooks(c *gin.Context){
    var books []Book
    db.Find(&books)
	c.JSON(200,books)
}

func getTopics(c *gin.Context){
    var topics []Topic
    db.Find(&topics)
	c.JSON(200,topics)
}

func createOrEditBook(c *gin.Context){
    var book Book
    c.BindJSON(&book)
    fmt.Println(book)
    book = createOrEditBookModel(book)
	c.JSON(200,book)
}

func getBook(c *gin.Context){
    id := c.Param("id")
    //db.First(&book,id)
    bookChapters := getBookModel(id)
    c.JSON(200,bookChapters)
}

func deleteBook(c *gin.Context){
    var book Book
    id := c.Param("id")
    book = deleteBookModel(id)
    c.JSON(200,book)
}

func createOrEditChapter(c *gin.Context){
    var chapter Chapter
    bookId := c.Param("id")
    c.BindJSON(&chapter)
    bookChapter := createOrEditChapterModel(bookId,chapter)
    c.JSON(200,bookChapter)
}


