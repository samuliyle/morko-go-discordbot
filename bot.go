package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/samuliyle/morko-go-discordbot/commands"
	"github.com/samuliyle/morko-go-discordbot/config"
	"github.com/samuliyle/morko-go-discordbot/database"
)

func main() {

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// log.SetOutput(file)

	dg, err := discordgo.New("Bot " + config.Secrets.Token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Println("Gracefully exiting...")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore bots
	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, config.Config.CommandPrefix) {
		database.LogMessage(s, m)
		return
	}

	commands.ParseCommand(s, m, func() string {
		return strings.TrimPrefix(m.Content, config.Config.CommandPrefix)
	}())
}
