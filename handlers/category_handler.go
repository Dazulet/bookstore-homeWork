package handlers

import (
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{
	{ID: 1, Name: "issecai"},
	{ID: 2, Name: "Fantasy"},
	{ID: 3, Name: "anime"},
	{ID: 4, Name: "Science"},
	{ID: 5, Name: "Horror"},
}

func GetCategory(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	category.ID = len(categories) + 1
	categories = append(categories, category)
	c.JSON(http.StatusCreated, category)
}
