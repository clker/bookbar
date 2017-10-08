package main

import (
        "github.com/gin-gonic/gin"
        "github.com/gin-contrib/cors"
)

func main() {

    open_db()

	r := gin.Default()
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:8080","http://127.0.0.1"}
    //config.AllowOrigins = []string{"*"}
    config.AllowMethods = []string{"GET", "PUT", "PATCH", "POST","DELETE"}
    //config.AllowHeaders = []string{"Origin"}
    //config.AllowCredentials = false
    //config.ExposeHeaders = []string{"Content-Length"}
    r.Use(cors.New(config))
    //r.Use(cors.Default())
    //r.Use()

	r.GET("/books", get_books)
	r.GET("/topics", get_topics)
    r.DELETE("/book/:id", delete_book)
	r.POST("/book", create_or_edit_book)
    r.GET("/book/:id", get_book)
    //r.POST("/book/:id", create_or_edit_book)
    r.Run(":8000") // listen and serve on 0.0.0.0:8080

    defer close_db()
}
