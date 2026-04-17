package main

import (
	"bookstore/config"
	"bookstore/def"
	"bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	r := gin.Default()

	r.POST("/login", handlers.Login)

	r.GET("/books", handlers.GetBook)
	r.GET("/books/:id", handlers.GetBookByID)

	protected := r.Group("/books")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/favorites", handlers.GetFavoriteBooks)
		protected.PUT("/:id/favorites", handlers.AddFavoriteBook)
		protected.DELETE("/:id/favorites", handlers.RemoveFavoriteBook)

		protected.POST("/", handlers.AddBook)
		protected.PUT("/:id", handlers.UpdateBook)
		protected.DELETE("/:id", handlers.DeleteBook)
	}

	r.Run(":8080")
}
