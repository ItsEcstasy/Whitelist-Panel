package config

import (
	"encoding/json"
	"log"
	"os"
)

type cfg struct {
	Token            string   `json:"token"`
	BotID            string   `json:"bot_id"`
	Prefix           string   `json:"prefix"`
	Brand            string   `json:"brand"`
	AuthenticatedIds []string `json:"authenticated_ids"`
	PrefixEnabled    bool     `json:"prefix_enabled"`
	SlashEnabled     bool     `json:"slash_enabled"`
	RemoveCommands   bool     `json:"remove_cmds"`
	GuildID          string   `json:"guild_id"`

	PaymentMethods struct {
		PayPal   string `json:"paypal"`
		CashApp1 string `json:"cashapp1"`
		CashApp2 string `json:"cashapp2"`
		CashApp3 string `json:"cashapp3"`
		ApplePay string `json:"applepay"`
		Venmo    string `json:"venmo"`
		BTC      string `json:"btc"`
		LTC      string `json:"ltc"`
		ETH      string `json:"eth"`
		BCH      string `json:"bch"`
		Card     string `json:"card"`
	} `json:"payment_methods"`
}

var (
	Config *cfg
)

func Load() {
	f, err := os.ReadFile("./assets/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = json.Unmarshal(f, &Config); err != nil {
		log.Fatal(err.Error())
	}
}
