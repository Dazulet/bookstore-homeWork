package handlers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavoriteBooks(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var favoriteBooks []models.FavoriteBook
	err := config.DB.Preload("Book").
		Where("user_id = ?", userID).
		Offset(offset).Limit(limit).
		Find(&favoriteBooks).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var books []models.Book
	for _, fb := range favoriteBooks {
		books = append(books, fb.Book)
	}

	c.JSON(http.StatusOK, books)
}

func AddFavoriteBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("userID").(uint)

	var book models.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	fav := models.FavoriteBook{
		UserID: userID,
		BookID: uint(bookID),
	}

	if err := config.DB.Create(&fav).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already in favorites or error"})
		return
	}

	c.JSON(http.StatusCreated, fav)
}

func RemoveFavoriteBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	userID := c.MustGet("userID").(uint)

	result := config.DB.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&models.FavoriteBook{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favorite record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Removed from favorites"})
}
