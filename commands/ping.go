package commands

import (
    dg "github.com/bwmarrin/discordgo"
    "fmt"
)

func Ping(s *dg.Session, m *dg.MessageCreate, args []string) {
    var ping string = fmt.Sprintf("%f", s.HeartbeatLatency().Seconds() * 1000) + "ms"
    s.ChannelMessageSend(m.ChannelID, "🏓 Pong! Heartbeat Latency: " + ping)
    return
}
