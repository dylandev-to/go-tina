package main

import (
	"fmt"
	"go-tina/internal/database"
	"go-tina/internal/discord"
	"go-tina/pkg/utils"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
		return
	}

	dg, err := discord.StartDiscord()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = database.StartDatabase(path.Join(utils.GetCwd(), "db", "bot.db"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}
