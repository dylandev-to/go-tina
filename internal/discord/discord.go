package discord

import (
	"fmt"
	"go-tina/internal/constants"
	"go-tina/pkg/events"

	"github.com/bwmarrin/discordgo"
)

func StartDiscord() (*discordgo.Session, error) {
	constants.LoadDiscordConfig()
	discordConstants := constants.GetDiscordConstants()

	if discordConstants.Token == "" {
		return nil, fmt.Errorf("token was not provided")
	}

	dg, err := discordgo.New("Bot " + discordConstants.Token)
	if err != nil {
		return nil, fmt.Errorf("couldn't log in")
	}

	events.HandleEvents(dg)

	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("error started listening")
	}

	return dg, nil
}
