package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Poll(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	question := strings.Join(args[1:], " ")

	pollMessage, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Poll: %s", question))
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error creating poll")
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)

	err = s.MessageReactionAdd(m.ChannelID, pollMessage.ID, "✅")
	if err != nil {
		fmt.Println("Error adding reaction:", err)
	}
	err = s.MessageReactionAdd(m.ChannelID, pollMessage.ID, "❌")
	if err != nil {
		fmt.Println("Error adding reaction:", err)
	}
}
func FakeBan(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(m.Mentions) < 1 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s <user>", args[0]))
		return
	}

	mentionedUser := m.Mentions[0]
	mentionedUserMention := mentionedUser.Mention()

	embed := NewEmbed().
		SetTitle("Ban Success").
		SetDescription(fmt.Sprintf("Successfully banned %s", mentionedUserMention)).
		SetColor(0xFF0000)

	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
	if err != nil {
		fmt.Println("Failed to send embed:", err)
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)

	guild, err := s.Guild(m.GuildID)
	if err != nil {
		return
	}
	guildName := guild.Name

	dmChannel, _ := s.UserChannelCreate(mentionedUser.ID)
	s.ChannelMessageSend(dmChannel.ID, fmt.Sprintf("You have been banned in: __%s__ by <@%s>", guildName, m.Author.ID))

	time.Sleep(5 * time.Second)
	embed.SetColor(0xFF0000).
		SetFooter("ITS A JOKE!!!", "https://avatars.githubusercontent.com/u/59181303?v=4")
	s.ChannelMessageEditEmbed(m.ChannelID, msg.ID, embed.MessageEmbed)
}
