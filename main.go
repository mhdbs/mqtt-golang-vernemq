package main

import (
	"fmt"
	routes "mhdbs/mqtt-golang-vernemq/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var configMqttHooksPort string

func main() {
	viper.SetConfigName("global")
	viper.AddConfigPath("./config")
	configError := viper.ReadInConfig()

	if configError != nil {
		fmt.Println("Fatal error config file: %s \n", configError)
		panic("Fatal error")
	}
	router := gin.Default()

	v1 := router.Group("/")
	routes.WebHooks(v1.Group("/"))
	configMqttHooksPort = viper.GetString("mqttHooks.port")
	PortEnv := os.Getenv("PORT")
	if len(PortEnv) > 0 {
		configMqttHooksPort = PortEnv
	}
	fmt.Println("Config: MqttHooks port: ", configMqttHooksPort)
	err := router.Run("0.0.0.0" + configMqttHooksPort)
	if err != nil {
		fmt.Println("Cannot start Server: %v", err)
	}
}
