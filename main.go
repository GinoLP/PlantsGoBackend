package main

import (
	"github.com/gin-gonic/gin"
	"plants/controllers"
	"plants/inits"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.AddPlant)
	r.GET("/", controllers.GetAllPlants)
	r.GET("/:id", controllers.GetPlant)
	r.DELETE("/:id", controllers.DeletePlant)
	err := r.Run()
	if err != nil {
		return
	}
}
