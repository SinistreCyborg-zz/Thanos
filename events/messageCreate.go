package events

import (
    "os"
    str "strings"
    dg "github.com/bwmarrin/discordgo"
)

var prefix string = os.Getenv("prefix")
func MessageCreate(s *dg.Session, m *dg.MessageCreate) {

    // Ignore messages that don't start with the prefix.
    if str.HasPrefix(m.Content, prefix) == false {
        return
    }

    // Useful when responding to commands.
    command := strings.Split(strings.TrimPrefix(m.Content, prefix), " ")[0]
    args := strings.Split(strings.TrimSpace(strings.TrimPrefix(m.Content, prefix + command)), " ")

}
