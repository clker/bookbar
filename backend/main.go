package main

import (
        "github.com/gin-gonic/gin"
        "github.com/gin-contrib/cors"
)

func main() {

    open_db()

	r := gin.Default()
	r.GET("/books", get_books)
	r.GET("/topics", get_topics)
	r.POST("/book", add_book)

    //config := cors.DefaultConfig()
    //config.AllowOrigins = []string{"http://localhost:8080"}
    //config.AllowMethods = []string{"GET", "PUT", "PATCH", "POST"}
    //config.AllowHeaders = []string{"Origin"}
    //config.AllowCredentials = false
    //config.ExposeHeaders = []string{"Content-Length"}
    //r.Use(cors.New(config))
    r.Use(cors.Default())
    r.Run(":8000") // listen and serve on 0.0.0.0:8080

    defer close_db()
}
