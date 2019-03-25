package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// opts := MQTT.NewClientOptions().AddBroker("ws://127.0.0.1:9001/mqtt") //for websocket
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883") //for tcp connection

	password := "test"
	opts.SetPassword(password)
	opts.SetUsername("test")
	opts.SetDefaultPublishHandler(f)
	topic := "dht11"

	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error", token)
		panic(token.Error())
	} else {
		fmt.Printf("Connected to server\n")
	}
	<-c
}
