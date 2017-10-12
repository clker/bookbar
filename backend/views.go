package main
import (
        "github.com/gin-gonic/gin"
        "fmt"
        "strconv"
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

//c.Param("id") is book_id;
//c.Param("ch_id") is the order number of the chapter
func getBook(c *gin.Context){
    var ch_order_id int
    var err error
    id := c.Param("id")
    ch_id := c.Param("ch_id")
    if(ch_id == ""){
        ch_order_id = 0
    }else{
        ch_order_id,err = strconv.Atoi(ch_id)
        if err != nil {
            c.AbortWithError(500,err)
        }
    }
    bookChapters := getBookModel(id,ch_order_id)
    //db.First(&book,id)
    c.JSON(200,bookChapters)
}

func deleteBook(c *gin.Context){
    var book Book
    id := c.Param("id")
    book = deleteBookModel(id)
    c.JSON(200,book)
}

func deleteChapter(c *gin.Context){
    var chapter Chapter
    id := c.Param("id")
    chapter = deleteChapterModel(id)
    c.JSON(200,chapter)
}

func createOrEditChapter(c *gin.Context){
    var chapter Chapter
    bookId := c.Param("id")
    c.BindJSON(&chapter)
    bookChapter := createOrEditChapterModel(bookId,chapter)
    c.JSON(200,bookChapter)
}


