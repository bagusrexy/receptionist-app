package router

import (
	"github.com/bagusrexy/test-dataon/controller"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	controllers := new(controller.GuestController)
	router.GET("/", controllers.HealthCheck)
	router.POST("/guests", controllers.RegisterGuest)
	router.PUT("/guests/:id/checkout", controllers.CheckOutGuest)
	router.POST("/guests/:id/photo", controllers.UploadPhoto)
}
