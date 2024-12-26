package gifs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var client = http.Client{}

const url = "https://api.otakugifs.xyz/gif?reaction=%s"

func GetGif(reaction string) (string, error) {
	resp, err := client.Get(fmt.Sprintf(url, reaction))
	if err != nil {
		return "", fmt.Errorf("error getting gif %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %v", err)
	}

	var data map[string]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", fmt.Errorf("error parsing data %v", err)
	}
	gifUrl, ok := data["url"]
	if !ok {
		return "", fmt.Errorf("url does not exist in data %v", err)
	}
	return gifUrl, nil
}
