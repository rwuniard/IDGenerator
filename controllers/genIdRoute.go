package controllers

import (
	"idgenerator/algorithms"

	"github.com/gin-gonic/gin"
)

func GenId(c *gin.Context) {
	id := algorithms.GenerateID(1, 1)
	c.JSON(200, gin.H{
		"id": id,
	})
}
