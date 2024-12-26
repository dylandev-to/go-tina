package actions

import (
	"fmt"
	"go-tina/internal/constants"
	"go-tina/pkg/gifs"
	"go-tina/pkg/utils"

	"github.com/bwmarrin/discordgo"
)

func Kiss(s *discordgo.Session, m *discordgo.MessageCreate) {
	react, err := gifs.GetGif("kiss")
	if err != nil {
		fmt.Println(err)
		return
	}

	mentions, err := utils.GetMentions(m.Content)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, constants.GetDiscordConstants().Config.Messages["mention_required"], m.Reference())
		return
	}

	user, err := s.GuildMember(m.GuildID, mentions[0])
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, constants.GetDiscordConstants().Config.Messages["mention_required"], m.Reference())
		return
	}

	s.ChannelMessageSendEmbedReply(m.ChannelID,
		&discordgo.MessageEmbed{
			Description: fmt.Sprintf(constants.GetDiscordConstants().Config.CmdMessages["kiss"], m.Member.Nick, user.Nick),
			Image: &discordgo.MessageEmbedImage{
				URL: react,
			},
		}, m.Reference())
}
