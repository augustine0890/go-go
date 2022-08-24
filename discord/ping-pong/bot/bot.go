package bot

import (
	"fmt"
	"ping-pong/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running... Press CTRL-C to exit.")
}

func messageHandler(s *discordgo.Session, m *discordgo.Message) {
	// if m.Author.ID == BotID {
	// return
	// }

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}
}
