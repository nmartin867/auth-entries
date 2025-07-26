package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nmartin867/auth-logs/initializers"
	"github.com/nmartin867/auth-logs/models"
)

func AuthEntryIndex(c *gin.Context) {
	// Get all the entries
	var entries []models.AuthEntry
	initializers.DB.Find(&entries)

	// Return entries in response
	c.JSON(200, gin.H{
		"entries": entries,
	})
}
