package commands

import (
	"Vanity/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Payments(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	embed := NewEmbed().
		SetTitle("Accepted Payment Methods (temp removing soon)").
		SetDescription("").
		AddField("ApplePay/AppleCash", fmt.Sprintf("Send to this email on iMessage:```%s```", config.Config.PaymentMethods.ApplePay)).
		AddField("CashApp", fmt.Sprintf("Send to this cashtag: ```%s```> if payment fails, send to **%s** or **%s**", config.Config.PaymentMethods.CashApp1, config.Config.PaymentMethods.CashApp2, config.Config.PaymentMethods.CashApp3)).
		AddField("Venmo", fmt.Sprintf("Send to this username: ```%s```", config.Config.PaymentMethods.Venmo)).
		AddField("PayPal send as friends & family (F&F)", fmt.Sprintf("```%s```", config.Config.PaymentMethods.PayPal)).
		AddField("BTC", fmt.Sprintf("```%s```", config.Config.PaymentMethods.BTC)).
		AddField("LTC", fmt.Sprintf("```%s```", config.Config.PaymentMethods.LTC)).
		AddField("ETH", fmt.Sprintf("```%s```", config.Config.PaymentMethods.ETH)).
		AddField("BCH", fmt.Sprintf("```%s```", config.Config.PaymentMethods.BCH)).
		AddField("Debit/Credit Card", fmt.Sprintf("```%s```", config.Config.PaymentMethods.Card)).
		SetColor(0x3498db).
		SetFooter("Love from Ecstasy", "https://avatars.githubusercontent.com/u/59181303?v=4").
		SetThumbnail("https://avatars.githubusercontent.com/u/59181303?v=4")

	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
}

func PurchaseKey(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	GamePass := "https://Example-whatever/"
	UserID := "1055337846657007648"

	embed := discordgo.MessageEmbed{
		Title:       "",
		Description: "temp description",
		Color:       0x00ff00,
		Fields: []*discordgo.MessageEmbedField{{
			Name:   " Link",
			Value:  fmt.Sprintf("[Click here](%s) example", GamePass),
			Inline: false,
		}, {
			Name:   "Contact",
			Value:  fmt.Sprintf("If you need help contact <@%s>", UserID),
			Inline: false,
		}}}

	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Embed: &embed,
	})
}
