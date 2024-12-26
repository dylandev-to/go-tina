package utils

import (
	"fmt"
	"regexp"
)

func GetMentions(text string) ([]string, error) {
	re := regexp.MustCompile(`<@(\d+)>`)
	matches := re.FindAllStringSubmatch(text, -1)
	var ids []string
	for _, match := range matches {
		ids = append(ids, match[1])
	}
	if len(ids) > 0 {
		return ids, nil
	}
	return nil, fmt.Errorf("mention required")
}

func GetDisplayName(nick string, username string) string {
	if nick != "" {
		return nick
	}
	return username
}
