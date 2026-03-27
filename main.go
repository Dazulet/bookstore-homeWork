package main

import (
	"bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", handlers.GetAuthors)
	r.POST("/users", handlers.AddAuthor)

	r.GET("/books", handlers.GetBook)
	r.POST("/books", handlers.AddBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.GET("/books/:id", handlers.GetBookByID)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.GET("/categories", handlers.GetCategory)
	r.POST("/categories", handlers.AddCategory)
}
