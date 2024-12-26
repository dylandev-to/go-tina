package main

import (
	"fmt"
	"go-tina/internal/database"
	"go-tina/internal/discord"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func nanosecondsToMilliseconds(nanoseconds int64) float64 {
	return float64(nanoseconds) / 1000000
}

func main() {
	start := time.Now()
	var count = 0
	for count < 1000000000000 {
		count++
	}
	elapsed := time.Since(start)
	fmt.Printf("%d nanoseconds is equal to %.3f milliseconds\n", elapsed.Nanoseconds(), nanosecondsToMilliseconds(elapsed.Nanoseconds()))
	return

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

	db, err := database.StartDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
	db.Close()
}
