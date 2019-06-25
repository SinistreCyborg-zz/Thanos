package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    dg "github.com/bwmarrin/discordgo"
    "ThanosCord/events"
)

func main() {

    // Create a new Discord session with the provided token.
    bot, err := dg.New("Bot " + os.Getenv("token"))
    if err != nil {
        log.Fatal("Error creating Discord session...", err)
        return
    }

    bot.AddHandler(events.MessageCreate)

    // Open Websocket connection.
    err = bot.Open()
    if err != nil {
        log.Fatal("Error opening connection...", err)
        return
    }

    // Wait here until Ctrl-C or some other term signal.
    log.Print("Bot is now running. Press Ctrl-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Close connection.
    bot.Close()

}
