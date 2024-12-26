package constants

import (
	"encoding/json"
	"fmt"
	"go-tina/pkg/utils"
	"os"
	"path/filepath"
)

var _discordConstants = &DiscordConstants{}

type DiscordConstants struct {
	ClientID string
	GuildID  string
	Token    string
	Config   DiscordConfig
}

type DiscordCommand struct {
	Names       []string `json:"names"`
	Description string   `json:"description"`
}

type DiscordConfig struct {
	Prefix      string                    `json:"prefix"`
	Messages    map[string]string         `json:"messages"`
	CmdMessages map[string]string         `json:"command_messages"`
	Commands    map[string]DiscordCommand `json:"commands"`
}

func LoadDiscordConfig() {
	_discordConstants.ClientID = os.Getenv("CLIENT_ID")
	_discordConstants.GuildID = os.Getenv("GUILD_ID")
	_discordConstants.Token = os.Getenv("TOKEN")

	file, err := os.Open(filepath.Join(utils.GetCwd(), "config", "discord.json"))
	if err != nil {
		fmt.Println("No discord.json config file found.", err)
		return
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&_discordConstants.Config); err != nil {
		fmt.Println("Couldn't parse discord.json:", err)
	}
}

func GetDiscordConstants() *DiscordConstants {
	return _discordConstants
}
