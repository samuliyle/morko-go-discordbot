package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("help", "Describes the usage of the command.", help).setHelp("args: [commandName]\n\nexample: !help help").add()
}

func help(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	if len(msglist) == 1 {
		return
	}
	commandName := msglist[1]
	command, ok := commands[commandName]
	if ok && commandName == strings.ToLower(command.Name) {
		commandDescription := "```Command: " + commandName + "\n" + command.Description
		if len(command.Help) != 0 {
			commandDescription += "\n\n" + command.Help
		}
		commandDescription += "```"
		s.ChannelMessageSend(m.ChannelID, commandDescription)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Command not found")
	}
}
