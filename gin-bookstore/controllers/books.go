package controllers

import (
	"net/http"

	"github.com/YashMovaliya/ginbookstore/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
