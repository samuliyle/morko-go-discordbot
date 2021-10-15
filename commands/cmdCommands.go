package commands

import (
	"sort"

	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("commands", "Lists commands", listCommands).add()
}

func listCommands(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	keys := make([]string, len(commands))
	i := 0
	for k := range commands {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	max := ""
	for _, value := range keys {
		if len(value) > len(max) {
			max = value
		}
	}
	maxLength := len(max)

	listOfCommands := "```" + padEnd("Command", maxLength-len("Command")) + " Description\n"
	for _, c := range keys {
		listOfCommands += "\n" + padEnd(c, maxLength-len(c)) + " " + commands[c].Description
	}
	listOfCommands += "```"
	s.ChannelMessageSend(m.ChannelID, listOfCommands)
}

func padEnd(str string, pad int) string {
	for i := 0; i <= pad; i++ {
		str += " "
	}
	return str
}
