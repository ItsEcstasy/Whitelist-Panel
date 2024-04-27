package commands

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Clear(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s clear <amount>", args[0]))
		return
	}

	numMessages, err := strconv.Atoi(args[1])
	if err != nil || numMessages <= 0 {
		s.ChannelMessageSend(m.ChannelID, "amount must be a positive number.")
		return
	}

	messages, _ := s.ChannelMessages(m.ChannelID, numMessages, "", "", "")
	for _, message := range messages {
		s.ChannelMessageDelete(m.ChannelID, message.ID)
	}
}
func Kick(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Usage: %s <user>", args[0]))
		return
	}

	userToKick := m.Mentions[0]
	var err error
	err = s.GuildMemberDelete(m.GuildID, userToKick.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error kicking user.")
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Successfully kicked ``%s#%s` (`%s`).", userToKick.Username, userToKick.Discriminator, userToKick.ID))
}

func Ban(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s <user>", args[0]))
		return
	}

	userToBan := m.Mentions[0]
	var err error
	err = s.GuildBanCreate(m.GuildID, userToBan.ID, 0)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error banning user.")
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Successfully banned `%s#%s` (`%s`).", userToBan.Username, userToBan.Discriminator, userToBan.ID))
}
