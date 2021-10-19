package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/samuliyle/morko-go-discordbot/config"

	"github.com/bwmarrin/discordgo"
)

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("mysql", config.Secrets.Database.Username+":"+config.Secrets.Database.Password+"@/"+config.Config.Database.Name)
	if err != nil {
		log.Println("error connecting to database, ", err)
		return
	}
	DB.SetConnMaxLifetime(0)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(2)
}

func LogMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if err != nil {
		return
	}

	_, err := DB.Query("INSERT INTO messages (message, userId, channelId, username, time) values(?, ?, ?, ?, NOW())", m.Content, m.Author.ID, m.ChannelID, m.Author.Username)
	if err != nil {
		log.Println("Failed to insert message, ", err)
	}
}
