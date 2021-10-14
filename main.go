package main 

import (
        "os"
	"os/signal"
	"syscall"
        "logs"
	"time"


	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.LOad(".env")
	if err != nil {
		panic("Unable .env file")
	}

	Token := os.Getenv("token")
	if Token == "" {
		panic("enter your bot token")
	}

	kk, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	kk.AddHandler(kuki.chatbot)
	
	err = kk.Open()
	if err != nil {
		log.Panic("connected to server")
	}

	log.Print("Bot is ready")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	kk.Close()
}
