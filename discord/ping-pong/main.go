package main

import (
	"fmt"
	"ping-pong/config"

	"ping-pong/bot"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return
}
