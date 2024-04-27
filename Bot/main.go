package main

import (
	"Vanity/bot"
	"Vanity/bot/commands"
	"Vanity/config"
)

func init() {
	config.Load()
	if config.Config.PrefixEnabled {
		commands.Load()
	}

}

func main() {
	bot.Start()
}
