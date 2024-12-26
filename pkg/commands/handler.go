package commands

import (
	"go-tina/internal/constants"
	"go-tina/pkg/commands/actions"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var discordConstants = constants.GetDiscordConstants()

type command struct {
	run func(*discordgo.Session, *discordgo.MessageCreate)
}

var commandsRun = map[string]*command{
	"ping": {
		run: ping,
	},
	"kiss": {
		run: actions.Kiss,
	},
}

func IsCommand(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	isComm := strings.HasPrefix(m.Content, discordConstants.Config.Prefix)
	if isComm {
		handle(s, m, getCommandName(strings.ToLower(m.Content)))
	}
	return isComm
}

func getCommandName(msg string) string {
	parts := strings.Fields(msg)
	if len(parts) > 0 {
		return strings.TrimPrefix(parts[0], discordConstants.Config.Prefix)
	}
	return ""
}

func getRun(cmdName string) *command {
	for key, cmd := range discordConstants.Config.Commands {
		if slices.Contains(cmd.Names, cmdName) {
			return commandsRun[key]
		}
	}
	return nil
}

func handle(s *discordgo.Session, m *discordgo.MessageCreate, cmdName string) {
	cmd := getRun(cmdName)
	if cmd == nil {
		s.ChannelMessageSend(m.ChannelID, discordConstants.Config.Messages["command_not_found"])
		return
	}
	cmd.run(s, m)
}
