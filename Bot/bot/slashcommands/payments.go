package slashcommands

import (
	"Vanity/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

/* Payment Methods | Change Crypto Addresses After Payments*/
func Payments(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := NewEmbed().
		SetTitle("Accepted Payment Methods").
		SetDescription("").
		AddField("ApplePay/AppleCash", config.Config.PaymentMethods.ApplePay).
		AddField("CashApp", fmt.Sprintf("Send to this cashtag: **%s**. If payment fails, send to **$TheGreenHat** or **$ItsJusNix2**", config.Config.PaymentMethods.CashApp1)).
		AddField("Venmo", fmt.Sprintf("Send to this username: ```@%s```", config.Config.PaymentMethods.Venmo)).
		AddField("BTC", fmt.Sprintf("```\n%s\n```", config.Config.PaymentMethods.BTC)).
		AddField("LTC", fmt.Sprintf("```\n%s\n```", config.Config.PaymentMethods.LTC)).
		AddField("ETH", fmt.Sprintf("```\n%s\n```", config.Config.PaymentMethods.ETH)).
		AddField("Debit/Credit Card", "```Coming soon```").
		SetColor(0x3498db).
		SetFooter("Love from Ecstasy", "https://avatars.githubusercontent.com/u/59181303?v=4").
		SetThumbnail("https://avatars.githubusercontent.com/u/59181303?v=4")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed.MessageEmbed,
			},
		},
	})
}
