package slashcommands

import (
	"Vanity/config"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

var RegisteredCommands []*discordgo.ApplicationCommand

type Command struct {
	ID          string
	Name        string
	Description string
	Type        discordgo.ApplicationCommandType
	Options     []*discordgo.ApplicationCommandOption
	Admin       bool
	Product     bool
	Execute     func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var (
	Commands = make(map[string]*Command)
	lock     sync.Mutex
)

var cmds = []Command{{
	// [+] ============== ECHO ============= [+]
	Name:        "echo",
	Description: "Repeat a Message",
	Type:        discordgo.ChatApplicationCommand,
	Admin:       false,
	Product:     false,
	Execute:     UserInfo,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "echo",
			Description: "Repeat a Message",
			Required:    false,
		},
	},
},
	// [+] ============== UTILITY ============= [+]
	{
		Name:        "userinfo",
		Description: "Display information about a user",
		Type:        discordgo.ChatApplicationCommand,
		Admin:       false,
		Product:     false,
		Execute:     UserInfo,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The user to get information about",
				Required:    true,
			},
		},
	},
	// [+] ============== KEY COMMANDS ============= [+]
	{
		Name:        "redeem",
		Description: "Redeem a License/Key",
		Type:        discordgo.ChatApplicationCommand,
		Admin:       false,
		Product:     false,
		Execute:     RedeemKey,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "key_license",
				Description: "Enter the Key/License to Redeem",
				Required:    true,
			},
		},
	}, {
		Name:        "blacklist",
		Description: "Blacklist a License/Key",
		Type:        discordgo.ChatApplicationCommand,
		Admin:       true,
		Product:     false,
		Execute:     BlacklistKey,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "key_license",
				Description: "Enter the Key/License to Blacklist",
				Required:    true,
			},
		},
	}, {
		Name:        "unblacklist",
		Description: "UnBlacklist a License/Key",
		Type:        discordgo.ChatApplicationCommand,
		Admin:       true,
		Product:     false,
		Execute:     UnBlacklistKey,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "key_license",
				Description: "Enter the Key/License to UnBlacklist",
				Required:    true,
			},
		},
	}, {
		Name:        "unlink",
		Description: "Unlink a Redeemed License/Key",
		Type:        discordgo.ChatApplicationCommand,
		Admin:       true,
		Product:     false,
		Execute:     UnlinkKey,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "key_license",
				Description: "Enter the Key/License to Unlink",
				Required:    true,
			},
		},
	}, {
		Name:        "lookup",
		Description: "Lookup a License/Key",
		Type:        discordgo.ChatApplicationCommand,
		Admin:       true,
		Product:     false,
		Execute:     LookupKey,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "key_license",
				Description: "Enter the Key/License to Lookup",
				Required:    true,
			},
		},
	},
}

func newCommand(c Command) {
	if _, ok := Commands[c.Name]; ok {
		log.Fatal("Conflicting Slash Commands")
	}
	lock.Lock()
	defer lock.Unlock()
	Commands[c.Name] = &c

	// registered commands
	RegisteredCommands = append(RegisteredCommands, &discordgo.ApplicationCommand{
		ID:          strconv.Itoa(int(time.Now().UnixNano())),
		Name:        c.Name,
		Type:        c.Type,
		Description: c.Description,
		Options:     c.Options,
	})
}

// func Load(s *discordgo.Session) {
// 	for _, v := range cmds {
// 		newCommand(v)
// 		if _, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
// 			ID:          strconv.Itoa(int(time.Now().UnixNano())),
// 			Name:        v.Name,
// 			Type:        v.Type,
// 			Description: v.Description,
// 			Options:     v.Options,
// 		}); err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}
// }

func Load(s *discordgo.Session) {
	for _, v := range cmds {
		newCommand(v)
	}

	for _, cmd := range RegisteredCommands {
		if _, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func Get(cmd string) (*Command, bool) {
	for _, v := range Commands {
		if strings.EqualFold(v.Name, cmd) {
			return v, true
		}
	}
	return nil, false
}

func HasPermission(interaction *discordgo.InteractionCreate) bool {
	for _, v := range config.Config.AuthenticatedIds {
		if v == interaction.Member.User.ID {
			return true
		}
	}
	return false
}
