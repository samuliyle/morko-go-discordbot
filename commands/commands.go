package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	commands = make(map[string]command)
)

type command struct {
	Name        string
	Description string
	Help        string

	Exec func(*discordgo.Session, *discordgo.MessageCreate, []string)
}

func ParseCommand(s *discordgo.Session, m *discordgo.MessageCreate, message string) {
	msglist := strings.Fields(message)

	if len(msglist) == 0 {
		return
	}

	commandName := strings.ToLower(msglist[0])

	if command, ok := commands[commandName]; ok && commandName == strings.ToLower(command.Name) {
		log.Printf("%v", msglist)
		command.Exec(s, m, msglist)
		return
	}
}

func (c command) add() command {
	commands[strings.ToLower(c.Name)] = c
	return c
}

func newCommand(name string, description string, f func(*discordgo.Session, *discordgo.MessageCreate, []string)) command {
	return command{
		Name:        name,
		Description: description,
		Exec:        f,
	}
}

func (c command) setHelp(help string) command {
	c.Help = help
	return c
}
