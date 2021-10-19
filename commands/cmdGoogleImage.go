package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/samuliyle/morko-go-discordbot/config"
)

type googleSearch struct {
	Items []item `json:"items"`
}

type item struct {
	Link  string `json:"link"`
	Title string `json:"title"`
}

func init() {
	disabled := false
	if len(config.Secrets.Google.Cx) == 0 || len(config.Secrets.Google.Id) == 0 {
		disabled = true
	}
	newCommand("google", "Fetches a Google image", googleImage).setDisabled(disabled).setHelp("args: [query]\n\nexample: !google forsen").add()
}

func googleImage(s *discordgo.Session, m *discordgo.MessageCreate, msglist []string) {
	if len(msglist) == 1 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	search := strings.Join(msglist[1:], "%20")
	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?q=%s&searchType=image&cx=%s&num=1&start=%d&imgsize=medium&key=%s&num=1", search, config.Secrets.Google.Cx, rand.Intn(10), config.Secrets.Google.Id)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "Failed to fetch Google image")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "Failed to fetch Google image")
		return
	}

	result := googleSearch{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "Failed to fetch Google image")
		return
	}
	if len(result.Items) == 0 {
		s.ChannelMessageSend(m.ChannelID, "No images found")
		return
	}

	image := result.Items[0]
	s.ChannelMessageSend(m.ChannelID, image.Link+" "+image.Title)
}
