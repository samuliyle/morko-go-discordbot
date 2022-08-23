package commands

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/samuliyle/morko-go-discordbot/database"
)

type message struct {
	message  string
	username string
	time     string
}

func init() {
	err := database.DB.Ping()
	if err != nil {
		log.Println("Database connection not available, disabling DB commands, ", err)
	}
	newCommand("quote", "Posts a random quote from the channel", quote).setHelp("args: (quote count)\ncount must be within 1-10\n\nexample: !quote 3").setDisabled(err != nil).add()
}

func quote(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	count := 1
	if len(msglist) > 1 {
		i, err := strconv.Atoi(msglist[1])
		if err == nil {
			count = i
		}
	}
	if count < 1 || count > 10 {
		count = 1
	}
	rows, err := database.DB.Query("CALL get_rands(?, ?)", count, m.ChannelID)
	if err != nil {
		log.Println("Failed to connect DB, ", err)
		return
	}
	quotes := ""
	for rows.Next() {
		var r message
		err := rows.Scan(&r.message, &r.username, &r.time)
		if err != nil {
			log.Println("Error casting row to struct", err)
			return
		}
		quotes += r.username + ": \"" + r.message + "\" (" + r.time + ")\n"
	}
	if len(quotes) == 0 {
		s.ChannelMessageSend(m.ChannelID, "This channel has no recorded messages")
		return
	}
	s.ChannelMessageSend(m.ChannelID, quotes)
}
