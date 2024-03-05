package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "Ardyn Rezky Fahreza",
		"umur": "21",
	})
}
