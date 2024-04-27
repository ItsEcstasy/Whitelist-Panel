package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func VerificationEmbed(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	TOSEmbed(s, m, "Terms Of Service", "**Product TOS**:\n- Methods will not be refunded due to the nature of the product.\n- Account Logs will not be refunded due to the nature of the product unless invalid (unless you opt for replacements).\n\n**I understand that**:\n- If I attempt to start drama, I can/will be kicked/banned.\n- If I am found DMing promos, I can/will be kicked/banned.\n\n**I accept that**:\n- When contacted by administrators, I am obligated to comply with their statements.\n- If I am kicked/banned for my own behavior before receiving a product, I will not receive a refund.\n- If I am found promoting I will be kicked/banned.\n- I recognize that if I am kicked/banned by an admin, I will not claim I've been scammed or ask for a refund.", 0x343aeb)
	VerifyEmbed(s, m)
}

func TOSEmbed(s *discordgo.Session, m *discordgo.MessageCreate, title, content string, color int) {
	embed := discordgo.MessageEmbed{
		Color:       color,
		Title:       title,
		Description: content,
		Thumbnail:   &discordgo.MessageEmbedThumbnail{},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

func VerifyEmbed(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := discordgo.MessageEmbed{
		Color:       0x343aeb,
		Title:       "Verification",
		Description: "[Click to verify](https://restorecord.com/verify/Link-Here)",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "By Verifying you agree to the terms of service.",
			IconURL: "https://images-ext-2.discordapp.net/external/dTuhTpecH0UvHpMyF5EYUBOUTqNETY5nNS5ee8D4hQ4/https/imgs.search.brave.com/_kKtAAe1p04vUnMjBBC2GhPBL4qplvkNODwtShKNDDM/rs%3Afit%3A500%3A500%3A1/g%3Ace/aHR0cHM6Ly9tZWRp/YS5naXBoeS5jb20v/bWVkaWEvbGxRTWpw/ZEN3amRyVkd6ejFk/L2dpcGh5LmdpZg.gif",
		},
	}
	s.ChannelMessageDelete(m.ChannelID, m.ID)
	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

func RulesEmbed(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	embed := NewEmbed().
		SetTitle("Rules / Guidelines").
		SetDescription("Please follow these rules in order to promote a friendly and safe environment for everyone.\n\n**Behavior**\nWe will not tolerate any form of racism, sexism, classism, homophobia or any other offensive language.\n\n**Content**\nDo not share any form of NSFW content (Pornography, Gore etc)\n\n**Channel Usage**\nPlease use each channel as intended. For example, using bot commands in the bot-commands channel.\n\n**Staff Rights**\nStaff will always have the right to issue instructions. If you break any rules, staff members can refer you to the according rule and decide on consequences.\n\n**Advertising**\nAdvertising a service un-related to Vanity is forbidden and can result in being banned from the server. If you ever receive a DM Advertisement from a non-authorized user, please let us know immediately in order to avoid potential scams. If you want to advertise something, contact someone with the @Admin/@Owner role\n\n**No Offensive Names & Profile Pictures**\nYou will be asked to change your name or profile picture if the staff deems them inappropriate.\n\n**Direct & Indirect Threats**\nThreatening to DDoS, Dox, Hack or any other form of malicious threat to another user is prohibited.\n\nVanity reserves the right to revoke your access to this Server for any violation of these rules.").
		SetColor(0x3498db).
		SetFooter("Your presence here implies acceptance of these rules, subject to future changes made without notice; it's your responsibility to stay updated.")

	s.ChannelMessageDelete(m.ChannelID, m.ID)
	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
}

func newAnnouncement(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s <content>", args[0]))
		return
	}

	content := strings.Join(args[1:], " ")

	embed := &discordgo.MessageEmbed{
		Color:       0xFFFF00,
		Title:       "🔔 Announcement",
		Description: content,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://media.tenor.com/imJm6NQvn_kAAAAC/announcement-batman.gif",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "",
			IconURL: "https://avatars.githubusercontent.com/u/59181303?v=4",
		},
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
