package main

import (
	"fmt"

	"github.com/go-git/go-git/v5/config"
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
