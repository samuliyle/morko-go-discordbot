package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("remind", "Sets a reminder", remind).setHelp("args: [minutes] (message)\n\nexample: !remind 15 food").add()
}

func remind(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	if len(msglist) == 1 {
		return
	}
	msg := ""
	if len(msglist) > 2 {
		msg = strings.Join(msglist[2:], " ")
	}
	i, err := strconv.Atoi(msglist[1])
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "Failed to parse time")
		return
	}
	if i < 0 {
		s.ChannelMessageSend(m.ChannelID, "Time must be at least 1 minute")
		return
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reminding you in %d minutes", i))
	time.AfterFunc(time.Duration(i)*time.Minute, func() {
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+fmt.Sprintf(" :fire: REMEMBER: \"%s\"! :fire:", msg))
	})
}
