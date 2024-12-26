package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Your bot %s is ready.\n", s.State.User.Username)
}
