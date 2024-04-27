package slashcommands

import (
	"github.com/bwmarrin/discordgo"
)

func Echo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	userMessage := options[0].Value.(string)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: userMessage,
		},
	})
}
