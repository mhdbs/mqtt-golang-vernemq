package controller

import (
	"fmt"
	"mhdbs/mqtt-golang-vernemq/model"

	"github.com/gin-gonic/gin"
)

//AuthOnRegister .
func AuthOnRegister(c *gin.Context) {
	fmt.Println("Auth on register")
	ok := model.Reply{"ok"}
	c.JSON(200, ok)

}

//AuthOnPublish .
func AuthOnPublish(c *gin.Context) {
	fmt.Println("auth on publish")
	ok := model.Reply{"ok"}
	c.JSON(200, ok)
}

//OnPublish .
func OnPublish(c *gin.Context) {
	fmt.Println("on publish")
	ok := model.Reply{"ok"}
	c.JSON(200, ok)
}

//AuthOnSubscribe .
func AuthOnSubscribe(c *gin.Context) {
	fmt.Println("auth on subscrbe")
	ok := model.Reply{"ok"}
	c.JSON(200, ok)
}
