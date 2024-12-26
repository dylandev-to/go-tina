package events

import (
	"github.com/bwmarrin/discordgo"
)

func HandleEvents(dg *discordgo.Session) {
	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
}
