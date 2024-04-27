package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RedeemKey(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// adding logic later
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Key Redeemed Successfully",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func BlacklistKey(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// adding logic later
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Key Blacklisted Successfully",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func UnBlacklistKey(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// adding logic later
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Key UnBlacklisted Successfully",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func UnlinkKey(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// adding logic later
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Key Unlinked Successfully",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
func LookupKey(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Adding logic later
	KeyInfo := map[string]interface{}{
		"User":     "<nil>",
		"Plan":     "<nil>",
		"Duration": "<nil>",
		"Active":   "<nil>",
		"Used":     "<nil>",
	}

	content := "Key Info:\n"
	for key, value := range KeyInfo {
		content += fmt.Sprintf("%s: %v\n", key, value)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
