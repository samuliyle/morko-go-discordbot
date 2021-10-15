package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	startTime = time.Now()
)

func init() {
	newCommand("uptime", "Bots uptime", uptime).add()
}

func uptime(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	duration := time.Since(startTime)
	s.ChannelMessageSend(m.ChannelID, duration.String())
}
