package events

import (
    "os"
    str "strings"
    dg "github.com/bwmarrin/discordgo"
    "ThanosCord/commands"
)

var prefix string = os.Getenv("prefix")
func MessageCreate(s *dg.Session, m *dg.MessageCreate) {

    // Ignore messages that don't start with the prefix.
    if str.HasPrefix(m.Content, prefix) == false {
        return
    }

    // Useful when responding to commands.
    command := str.Split(str.TrimPrefix(m.Content, prefix), " ")[0]
    args := str.Split(str.TrimSpace(str.TrimPrefix(m.Content, prefix + command)), " ")

    if command == "ping" {
        commands.Ping(s, m, args)
        return
    }

    if command == "snap" {
        commands.Snap(s, m, args)
    }

}
