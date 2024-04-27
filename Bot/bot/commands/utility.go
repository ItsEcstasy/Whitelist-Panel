package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ListRoles(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	guildRoles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed to retrieve server roles.")
		return
	}

	if len(guildRoles) == 0 {
		s.ChannelMessageSend(m.ChannelID, "No roles found in this server.")
		return
	}

	var roleList string
	for _, role := range guildRoles {
		roleList += fmt.Sprintf("- <@&%s> (ID: %s)\n", role.ID, role.ID)
	}

	embed := NewEmbed().
		SetTitle("Server Roles").
		SetDescription(roleList).
		SetColor(0x3498db).
		SetFooter("Requested by " + m.Author.Username).
		MessageEmbed

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
