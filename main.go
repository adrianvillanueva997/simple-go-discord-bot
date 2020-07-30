package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	prefix string = "!"
)

func initEnv() {
	err := gotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Println("Environment file loaded successfully")

}

func main() {
	initEnv()
	bot, err := discordgo.New("Bot " + os.Getenv("token"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.AddHandler(MessageHandler)
	// bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = bot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println("Bot running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = bot.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Bot logged out of discord")
}
func MessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	log.Println(message.Content, message.Author, message.ChannelID)
	if message.Author.ID == session.State.User.ID {
		return
	}
	if string(message.Content[0]) == prefix {
		splittedMessage := strings.Split(message.Content, " ")
		fmt.Println(splittedMessage)
		switch splittedMessage[0] {
		case "!cat":
			photoUrl := getCatPhoto()
			_, _ = session.ChannelMessageSend(message.ChannelID, photoUrl)
			break
		default:

			break

		}

	} else {
		if strings.Contains(message.Content, "cbt") {
			_, _ = session.ChannelMessageSend(message.ChannelID, "Cock and ball torture (CBT), penis torture or dick torture is a sexual activity involving application of pain or constriction to the penis or testicles. This may involve directly painful activities, such as genital piercing, wax play, genital spanking, squeezing, ball-busting, genital flogging, urethral play, tickle torture, erotic electrostimulation, kneeing or kicking!")

		}
		if message.Content == "hello there" {
			_, _ = session.ChannelMessageSend(message.ChannelID, "General Kenobi!")
		}
	}

}
