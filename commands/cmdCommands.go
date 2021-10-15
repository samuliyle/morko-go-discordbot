package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("commands", "Lists commands", listCommands).add()
}

func listCommands(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	listOfCommands := "```"
	for _, c := range commands {
		listOfCommands += "\n" + c.Name + ": " + c.Description
	}
	listOfCommands += "```"
	s.ChannelMessageSend(m.ChannelID, listOfCommands)
}
