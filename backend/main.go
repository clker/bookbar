package main

import (
        "github.com/gin-gonic/gin"
        "github.com/gin-contrib/cors"
)

func main() {

    openDb()

	r := gin.Default()
    config := cors.DefaultConfig()
    //config.AllowAllOrigins = true
    config.AllowOrigins = []string{"http://localhost:8080","http://127.0.0.1:8080"}
    //config.AllowOrigins = []string{"*"}
    config.AllowMethods = []string{"GET", "PUT", "PATCH", "POST","DELETE"}
    //config.AllowHeaders = []string{"Origin"}
    //config.AllowCredentials = false
    //config.ExposeHeaders = []string{"Content-Length"}
    r.Use(cors.New(config))
    //r.Use(cors.Default())
    //r.Use()
    //gin.DisableConsoleColor()
    //f, _ := os.Create("gin.log")
    //gin.DefaultWriter = io.MultiWriter(f)

	r.GET("/books", getBooks)
	r.GET("/topics", getTopics)
    r.DELETE("/book/:id", deleteBook)
    r.DELETE("/ch/:id", deleteChapter)
	r.POST("/book", createOrEditBook)
    r.GET("/book/:id", getBook)
    r.POST("/book/:id/chapter", createOrEditChapter)
    r.GET("/book/:id/ch/:ch_id", getBook)
    r.Run(":8000") // listen and serve on 0.0.0.0:8080

    defer closeDb()
}
