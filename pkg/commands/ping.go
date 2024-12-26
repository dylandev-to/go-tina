package commands

import (
	"go-tina/internal/constants"

	"github.com/bwmarrin/discordgo"
)

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, constants.GetDiscordConstants().Config.CmdMessages["ping"], m.Reference())
}
