package commands

import (
    dg "github.com/bwmarrin/discordgo"
    "log"
    "math"
)

func Snap(s *dg.Session, m *dg.MessageCreate, args []string) {

    // Get list of members
    members, err := s.GuildMembers(m.GuildID, "", 1000)
    if err != nil {
        log.Println(err)
        return
    }

    if err != nil {
        log.Println(err)
        return
    }

    for i, member := range members {
        if i % 2 == 0 {
            s.GuildMemberDeleteWithReason(m.GuildID, member.User.ID, "Thanos wills it.")
        }
    }

    s.ChannelMessageSend(m.ChannelID, "Banned half the members. 💛❤️💙💜🧡")

}
