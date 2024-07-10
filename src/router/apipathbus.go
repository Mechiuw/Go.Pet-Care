package router

import (
	"ginpet/src/controller"

	"github.com/gin-gonic/gin"
)

func API_BUS() {
	router_bus := gin.Default()
	api := router_bus.Group("/api")
	{
		client := api.Group("/client")
		{
			client.POST("/", controller.Create)
			client.PUT("/:id", controller.Update)
			client.DELETE("/:id", controller.Delete)
			client.GET("/all", controller.GetAll)
			client.GET("/:id", controller.GetById)
		}
		pet := api.Group("/pet")
		{
			pet.POST("/", controller.CreatePet)
			pet.PUT("/:id", controller.UpdatePet)
			pet.DELETE("/:id", controller.DeletePet)
			pet.GET("/all", controller.GetAllPet)
			pet.GET("/:id", controller.GetByIdPet)
		}
		sp := api.Group("/service-provider")
		{
			sp.POST("/", controller.CreateSp)
			sp.PUT("/:id", controller.UpdateSp)
			sp.DELETE("/:id", controller.DeleteSp)
			sp.GET("/all", controller.GetAllSp)
			sp.GET("/:id", controller.GetByIdSp)
		}
		app := api.Group("/appointment")
		{
			app.POST("/", controller.CreateAppointment)
			app.PUT("/:id", controller.UpdateAppointment)
			app.DELETE("/:id", controller.)
			app.GET("/all", controller.)
			app.GET("/:id", controller.)
		}
		review := api.Group("/review")
		{
			review.POST("/", controller.)
			review.PUT("/:id", controller.)
			review.DELETE("/:id", controller.)
			review.GET("/all", controller.)
			review.GET("/:id", controller.)
		}
	}

	router_bus.Run(":8080")
}
