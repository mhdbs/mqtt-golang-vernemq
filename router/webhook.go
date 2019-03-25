package routes

import (
	"mhdbs/mqtt-golang-vernemq/controller"

	"github.com/gin-gonic/gin"
)

//WebHooks .
func WebHooks(router *gin.RouterGroup) {
	router.POST("/auth_on_register", controller.AuthOnRegister)
	router.POST("/auth_on_publish", controller.AuthOnPublish)
	router.POST("/auth_on_subscribe", controller.AuthOnSubscribe)
	router.POST("/on_publish", controller.OnPublish)
}
