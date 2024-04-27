// Portal: https://discord.com/developers/applications/1154580863544209408/oauth2/general
/*

To do:
	- addemoji (name, attachment)
	- add voice commands
	- cleanup

*/

package commands

import (
	"Vanity/config"
	"log"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// Command represents a command structure
type Command struct {
	Name        string
	Alias       []string
	Description string
	AdminOnly   bool
	Menu        bool
	Misc        bool
	Util        bool
	Setting     bool
	Product     bool
	Execute     func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

var (
	Commands = make(map[string]*Command)
	lock     sync.Mutex
	cmds     = []Command{
		// [+] =============== SETTINGS =============== [+]
		{
			Name:        "invite",
			Alias:       []string{"Invite", "INVITE"},
			Description: "Invite to a server",
			AdminOnly:   false,
			Menu:        false,
			Misc:        false,
			Util:        false,
			Product:     false,
			Setting:     true,
			Execute:     Invite,
		}, {
			Name:        "shutdown",
			Alias:       []string{"Shutdown", "SHUTDOWN"},
			Description: "Shutdown the bot",
			AdminOnly:   true,
			Menu:        false,
			Misc:        false,
			Util:        false,
			Product:     false,
			Setting:     true,
			Execute:     Shutdown,
		}, {
			Name:        "config",
			Alias:       []string{"Config", "CONFIG"},
			Description: "Check the configuration settings",
			AdminOnly:   true,
			Menu:        false,
			Misc:        false,
			Util:        false,
			Product:     false,
			Setting:     true,
			Execute:     CheckConfig,
		},
		// [+] =============== MODERATION =============== [+]
		{
			Name:        "clear",
			Alias:       []string{"purge", "clean", "c"},
			Description: "Clear messages",
			AdminOnly:   true,
			Menu:        false,
			Misc:        false,
			Util:        false,
			Product:     false,
			Setting:     false,
			Execute:     Clear,
		},
		// [+] =============== MISC =============== [+]
		{
			Name:        "poll",
			Alias:       []string{"Poll", "POLL"},
			Description: "create a poll",
			AdminOnly:   false,
			Menu:        false,
			Misc:        true,
			Util:        false,
			Product:     false,
			Setting:     false,
			Execute:     Poll,
		}, {
			Name:        "bean",
			Alias:       []string{"Bean", "BEAN"},
			Description: "bean a user",
			AdminOnly:   false,
			Menu:        false,
			Misc:        true,
			Util:        false,
			Product:     false,
			Setting:     false,
			Execute:     FakeBan,
		},
		// [+] =============== UTILITY =============== [+]
		{
			Name:        "listroles",
			Alias:       []string{"showrole", "roles", "checkroles", "viewroles"},
			Description: "List all roles",
			AdminOnly:   true,
			Menu:        false,
			Misc:        false,
			Util:        true,
			Product:     false,
			Setting:     false,
			Execute:     ListRoles,
		},
		// [+] =============== SERVER =============== [+]
		{
			Name:        "verifyembed",
			Alias:       []string{"verify", "verification", "verifyembed"},
			Description: "Send the verification embed",
			AdminOnly:   true,
			Menu:        false,
			Misc:        false,
			Util:        false,
			Product:     false,
			Setting:     false,
			Execute:     VerificationEmbed,
		}, {
			Name:        "rulesembed",
			Alias:       []string{"rules"},
			Description: "Send the rules embed",
			AdminOnly:   true,
			Menu:        false,
			Misc:        false,
			Util:        false,
			Product:     false,
			Setting:     false,
			Execute:     RulesEmbed,
		}, {
			Name:        "announce",
			Alias:       []string{"news"},
			Description: "Create an announcement embed",
			AdminOnly:   true,
			Menu:        false,
			Misc:        true,
			Util:        false,
			Product:     false,
			Setting:     false,
			Execute:     newAnnouncement,
		},
		// [+] =============== PRODUCTS =============== [+]
		{
			Name:        "payments",
			Alias:       []string{"accepted"},
			Description: "Send list of payments accepted",
			AdminOnly:   false,
			Menu:        false,
			Misc:        false,
			Util:        false,
			Product:     false,
			Setting:     false,
			Execute:     Payments,
		}, {
			Name:        "purchase",
			Alias:       []string{"Purchase", "PURCHASE", "buy", "Buy", "BUY"},
			Description: "Purchase a key",
			AdminOnly:   false,
			Menu:        false,
			Misc:        false,
			Util:        false,

			Product: true,
			Setting: false,
			Execute: PurchaseKey,
		},
	}
)

func Load() {

	for _, cmd := range cmds {
		newCommand(cmd)
	}
}

func newCommand(c Command) {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := Commands[c.Name]; ok {
		// conflicting command names
		log.Fatalf("\033[38;5;207m[\033[97mError\033[38;5;207m]\033[93m conflicting command names\033[0m")
	}

	Commands[c.Name] = &c
}

func GetCommand(name string, m *discordgo.MessageCreate) (bool, *Command) {
	lock.Lock()
	defer lock.Unlock()

	// First, try to find a command with the given name
	cmd, ok := Commands[name]
	if ok && cmd != nil {
		// Check if the user is authorized to use the command
		if cmd.AdminOnly {
			for _, v := range config.Config.AuthenticatedIds {
				if strings.EqualFold(m.Author.ID, v) {
					return true, cmd
				}
			}
			// User is not an admin
			return false, cmd
		}
		return true, cmd
	}
	for _, cmd := range Commands {
		for _, alias := range cmd.Alias {
			if strings.EqualFold(name, alias) {
				if cmd.AdminOnly {
					for _, v := range config.Config.AuthenticatedIds {
						if strings.EqualFold(m.Author.ID, v) {
							return true, cmd
						}
					}
					return false, cmd
				}
				return true, cmd
			}
		}
	}

	return false, nil
}

func Activity(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "[🟢] **Rebooted Successfully**")
}
