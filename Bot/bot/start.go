package bot

import (
	"Vanity/bot/commands"
	"Vanity/bot/slashcommands"
	"Vanity/config"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	StartTime = time.Now()
)

func Start() {

	vanity, err := discordgo.New("Bot " + config.Config.Token)
	if err != nil {
		log.Fatal("\033[38;5;207m[\033[97mError\033[38;5;207m]\033[93m cannot open connection to Discord\033[38;5;207m: \033[97m", err)
		return
	}

	if config.Config.PrefixEnabled {
		vanity.AddHandler(messageCreate)
	}

	vanity.AddHandler(ready)
	if config.Config.SlashEnabled {
		vanity.AddHandler(handler)
	}
	err = vanity.Open()
	if err != nil {
		log.Fatal("\033[38;5;207m[\033[97mSession\033[38;5;207m] \033[38;5;98m \033[38;5;48mReady", err)
		return
	}

	defer vanity.Close()

	if config.Config.SlashEnabled {
		slashcommands.Load(vanity)
	}

	fmt.Println("[Bot] Running")

	// We want to wait for a termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc
}

func ready(session *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("\033[38;5;207m[\033[97mSession\033[38;5;207m] \033[92mRunning\033[0m")
	fmt.Printf("\033[38;5;207m[\033[97mBrand\033[38;5;207m] \033[92m%s\033[0m\n", config.Config.Brand)
	fmt.Printf("\033[38;5;207m[\033[97mUser\033[38;5;207m] \033[97m\033[92m%s \033[95m(\033[97m%s\033[95m)\033[0m\n", session.State.User.Username, session.State.User.ID)

	// mode
	if config.Config.SlashEnabled && config.Config.PrefixEnabled {
		fmt.Println("\033[38;5;207m[\033[97mMode\033[38;5;207m] \033[92mPrefix/Slash\033[0m\r\n")
	} else if config.Config.SlashEnabled {
		fmt.Println("\033[38;5;207m[\033[97mMode\033[38;5;207m] \033[92mSlash\033[0m")
	} else if config.Config.PrefixEnabled {
		fmt.Println("\033[38;5;207m[\033[97mMode\033[38;5;207m] \033[92mPrefix\033[0m")
	} else {
		fmt.Println("\033[38;5;207m[\033[97mMode\033[38;5;207m] \033[91mNot Running\033[0m")

		if config.Config.GuildID != "" {
			fmt.Printf("\033[38;5;207m[\033[97mGlobal\033[38;5;207m] \033[92m%s\033[0m", config.Config.GuildID)
		} else {
			fmt.Printf("\033[38;5;207m[\033[97mGlobal\033[38;5;207m] \033[91m%s\033[0m", config.Config.GuildID)
		}

		if config.Config.RemoveCommands {
			fmt.Printf("\n\033[38;5;207m[\033[97mRemove Slash On Close\033[38;5;207m] \033[92m%t\033[0m", config.Config.RemoveCommands)
		} else {
			fmt.Printf("\n\033[38;5;207m[\033[97mRemove Slash On Close\033[38;5;207m] \033[91m%t\033[0m\n", config.Config.RemoveCommands)
		}
	}

	fmt.Println("[Bot] Adding Slash Commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(slashcommands.RegisteredCommands))
	for i, v := range slashcommands.RegisteredCommands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, config.Config.GuildID, v)
		if err != nil {
			log.Panicf("\x1b[97mCannot create \x1b[95m'\x1b[97m%v\x1b[95m' \x1b[97mcommand\x1b[95m: \x1b[97m%v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("\x1b[97mPress Ctrl\x1b[95m+\x1b[97mC to exit")
	<-stop

	if config.Config.RemoveCommands {
		// fmt.Println("[Bot] Removing Slash Commands...")
		for _, v := range registeredCommands {
			err := session.ApplicationCommandDelete(session.State.User.ID, config.Config.GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
	fmt.Println("[Bot] Shutting Down")
}

func messageCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == session.State.User.ID {
		// fmt.Println(session.State.User.ID)
		return
	}

	args := strings.Split(m.Content, " ")

	if strings.HasPrefix(args[0], config.Config.Prefix) {
		stripped := strings.TrimPrefix(args[0], config.Config.Prefix)
		ok, command := commands.GetCommand(stripped, m)
		if !ok && command != nil {
			session.ChannelMessageSend(m.ChannelID, "You are not authorized to use this command.")
			return
		} else if ok {
			command.Execute(session, m, args)
		} else if !ok && command == nil {
			session.ChannelMessageSend(m.ChannelID, "Could not find command!")
		}
	}
}

func handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	command, ok := slashcommands.Get(data.Name)
	if ok && command != nil {
		if command.Admin {
			if slashcommands.HasPermission(i) {
				command.Execute(s, i)
			} else {
				s.ChannelMessageSend(i.ChannelID, "You do not have permission to use this command.")
			}
		} else {
			command.Execute(s, i)
		}
	}
}
