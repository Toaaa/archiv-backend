package controllers

import (
	"net/http"

	"github.com/Toaaa/archiv-backend/pkg/database"
	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	db, _ := database.DB.DB()

	if err := db.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "no connection to database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Ok",
	})
}
