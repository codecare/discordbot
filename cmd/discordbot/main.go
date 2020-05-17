package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
import "github.com/bwmarrin/discordgo"

func main() {


	if len(os.Args) < 2 {
		panic("missing filepath of file to send as argument")
	}
	token := os.Getenv("DISCORD_TOKEN")
	if len(token) < 2 {
		panic("env DISCORD_TOKEN is not set")
	}
	channelId := os.Getenv("DISCORD_CHANNEL_ID")
	if len(channelId) < 2 {
		panic("env DISCORD_CHANNEL_ID is not set")
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("session: ", err)
		return
	}

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer closeDiscord(discord)

	_, err = discord.ChannelMessageSend(channelId, string(content))
	if err != nil {
		log.Fatal("channel send", err)
	}
}

func closeDiscord(discord *discordgo.Session) {
	// Cleanly close down the Discord session.
	err := discord.Close()
	if err != nil {
		fmt.Println("error closing connection,", err)
		return
	}
}