package controllers

import (
	"idgenerator/algorithms"
	"log"

	"github.com/gin-gonic/gin"
)

type IdRequest struct {
	DataCenterId int `json:"DataCenterId"`
	MachineId    int `json:"MachineId"`
}

func GenId(c *gin.Context) {
	var idRequest IdRequest
	c.BindJSON(&idRequest)
	log.Println("dataCenterId:", idRequest.DataCenterId)
	log.Println("machineId:", idRequest.MachineId)
	id := algorithms.GenerateID(int(idRequest.DataCenterId), int(idRequest.MachineId))
	if id < 0 {
		c.JSON(400, gin.H{"message": "Invalid dataCenterId or machineId (must be between 0 - 31"})
		return
	}
	c.JSON(200, gin.H{
		"id": id,
	})
}
