package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"ping-pong/internal/config"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const defaultConfig = "./config/config.json"

var flagConfig = flag.String("c", defaultConfig, "The location of the config file.")

func main() {
	cfg, err := config.ReadConfig(*flagConfig)
	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	if err = s.Open(); err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	s.Close()
}
