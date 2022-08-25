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
		fmt.Println("Error creating Discord session,", err)
		panic(err)
	}

	// Register the messageCreate func as a callback for MessageCreate events
	s.AddHandler(messageCreate)

	if err = s.Open(); err != nil {
		fmt.Println("Error opening connection,", err)
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	s.Close()
}

// This function will be called (due to AddHandler above) every time a new message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping!")
	}
}
