package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("hello", "Prints hello", hello).add()
}

func hello(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	s.ChannelMessageSend(m.ChannelID, "Hello!")
}
