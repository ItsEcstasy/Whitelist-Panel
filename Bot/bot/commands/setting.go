package commands

import (
	"Vanity/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Shutdown(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "Shutting down...")
	fmt.Println("\033[38;5;207m[\033[97mSocket\033[38;5;207m] \033[91minActive (Restart Required)\033[0m")
	s.Close()
}

func Invite(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	userID := m.Author.ID
	dmChannel, err := s.UserChannelCreate(userID)
	if err != nil {
		return
	}

	BotID := config.Config.BotID
	if BotID == "" {
		fmt.Println("Error: The bot ID is not set in assets/config.json")
		return
	}

	inviteLink := fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=%s&scope=bot", BotID)

	embed := NewEmbed().
		SetTitle("Bot Invite Link").
		SetDescription(fmt.Sprintf("CLick [here] to invite me to a server(%s)", inviteLink)).
		SetColor(0x3498db).
		SetFooter("Love from Ecstasy", "https://avatars.githubusercontent.com/u/59181303?v=4")

	s.ChannelMessageSendEmbed(dmChannel.ID, embed.MessageEmbed)
}
func CheckConfig(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	authenticatedIDs := config.Config.AuthenticatedIds
	Admins := ""
	for _, userID := range authenticatedIDs {
		Admins += "<@" + userID + "> "
	}

	embed := NewEmbed().
		SetTitle("Configuration").
		SetDescription("").
		AddField("Command Prefix", fmt.Sprintf("```%s```", config.Config.Prefix)).
		AddField("Brand", fmt.Sprintf("```%s```", config.Config.Brand)).
		// AddField("Remove Slash Commands On Close", fmt.Sprintf("```%v```", config.Config.RemoveCommands)).
		// AddField("GuildID", fmt.Sprintf("```%s```", config.Config.GuildID)).
		AddField("Authenticated Users", Admins).
		AddField("Prefix Commands", fmt.Sprintf("```%t```", config.Config.PrefixEnabled)).
		AddField("Slash Commands", fmt.Sprintf("```%t```", config.Config.SlashEnabled)).
		SetColor(0x3498db).
		SetFooter("Love from Ecstasy", "https://avatars.githubusercontent.com/u/59181303?v=4").
		SetThumbnail("https://avatars.githubusercontent.com/u/59181303?v=4")

	s.ChannelMessageDelete(m.ChannelID, m.ID)
	s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
}
