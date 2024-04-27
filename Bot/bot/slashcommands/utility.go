package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func UserInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	var user *discordgo.User

	if len(options) > 0 && options[0].Type == discordgo.ApplicationCommandOptionUser {
		user = options[0].UserValue(s)
	} else {
		user = i.Member.User
	}

	member, err := s.GuildMember(i.GuildID, user.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to fetch user information",
			},
		})
		return
	}

	avatarURL := user.AvatarURL("")
	embed := &discordgo.MessageEmbed{
		Title:       "User Information",
		Description: "",
		Color:       0x3498db, // Blue color
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: avatarURL,
		},
		Fields: []*discordgo.MessageEmbedField{
			{Name: "User ID", Value: user.ID, Inline: true},
			{Name: "Username", Value: fmt.Sprintf("%s#%s", user.Username, user.Discriminator), Inline: true},
			{Name: "Accent Color", Value: fmt.Sprintf("%v", user.AccentColor), Inline: true},
			{Name: "Mention", Value: user.Mention(), Inline: true},
			{Name: "Is Bot", Value: fmt.Sprintf("%v", user.Bot), Inline: true},
			{Name: "Flags", Value: fmt.Sprintf("%d", user.Flags), Inline: true},
			{Name: "MFA Enabled", Value: fmt.Sprintf("%v", user.MFAEnabled), Inline: true},
			{Name: "Premium Type", Value: fmt.Sprintf("%d", user.PremiumType), Inline: true},
			{Name: "Public Flags", Value: fmt.Sprintf("%d", user.PublicFlags), Inline: true},
			{Name: "Is a Discord Official", Value: fmt.Sprintf("%v", user.System), Inline: true},
			{Name: "Verified User", Value: fmt.Sprintf("%v", user.Verified), Inline: true},
			{Name: "Joined", Value: member.JoinedAt.Format("2006-01-02 15:04:05"), Inline: true},
			{Name: "Permissions", Value: fmt.Sprintf("%v", member.Permissions), Inline: true},
			{Name: "Is Deafened", Value: fmt.Sprintf("%v", member.Deaf), Inline: true},
			{Name: "Is Pending", Value: fmt.Sprintf("%v", member.Pending), Inline: true},
			{Name: "Member User", Value: fmt.Sprintf("%v", member.User), Inline: true},
			{Name: "User Roles", Value: getRolesString(s, i.GuildID, member.Roles), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Love From Ecstasy",
			IconURL: "https://avatars.githubusercontent.com/u/59181303?v=4",
		},
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

func getRolesString(s *discordgo.Session, guildID string, roles []string) string {
	var rolesString string
	for _, roleID := range roles {
		rolesString += fmt.Sprintf("<@&%s> ", roleID)
	}
	return rolesString
}
