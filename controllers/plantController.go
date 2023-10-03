package controllers

import (
	"github.com/gin-gonic/gin"
	"plants/inits"
	"plants/models"
)

func AddPlant(ctx *gin.Context) {
	var plant struct {
		LatinName string
		Name      string
	}

	BindErr := ctx.BindJSON(&plant)
	if BindErr != nil {
		ctx.JSON(500, gin.H{"error": BindErr})
		return
	}

	post := models.Plant{Name: plant.Name, LatinName: plant.LatinName}
	returnedPlant, err := inits.InsertIntoColletion("Plants", post)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"data": returnedPlant})
}

func GetAllPlants(ctx *gin.Context) {
	plants, err := inits.GetAllInCollection("Plants")

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
	}

	ctx.JSON(200, gin.H{"data": plants})
}

func GetPlant(ctx *gin.Context) {
	id := ctx.Param("id")
	plant, err := inits.GetItemInCollectionWithId("Plants", id)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"data": plant})
}

func DeletePlant(ctx *gin.Context) {
	id := ctx.Param("id")

	nbDeleted, err := inits.DeleteItemInCollection("Plants", id)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}
	if nbDeleted == 0 {
		ctx.JSON(500, gin.H{"error": "no plant have been deleted"})
		return
	}

	ctx.JSON(200, gin.H{"data": "plant has been deleted successfully"})
}
