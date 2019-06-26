package commands

import (
    dg "github.com/bwmarrin/discordgo"
    "log"
)

func Snap(s *dg.Session, m *dg.MessageCreate, args []string) {

    // Get list of members
    members := &([]*dg.Member{})
    recurseMembers(members, GuildID, 0)

    // Get half of members
    var removal []dg.Member
    for index, member := range members {
        if index % 2 == 0 {
            removal = append(removal, member)
        }
    }

    // Ban them.
    for _, member := range removal {

    }

}

// Discord limits the members at 1000
func recurseMembers(memstore *[]*dg.Member, guildID, after int64) {

    members, err := dg.GuildMembers(guildID, after, 1000)
    if err != nil {
        log.Println(err)
        return
    }

    if len(members) == 1000 {
        recurseMembers(memstore, guildID, members[999].User.ID)
    }

    *memstore = append(*memstore, members...)
    return

}
