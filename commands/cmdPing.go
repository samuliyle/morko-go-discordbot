package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("ping", "Pings", ping).add()
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	s.ChannelMessageSend(m.ChannelID, "pong")
}
