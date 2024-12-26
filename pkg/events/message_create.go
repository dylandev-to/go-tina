package events

import (
	"go-tina/pkg/commands"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	commands.IsCommand(s, m)
}
