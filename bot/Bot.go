package bot

import (
	"fmt"
	"tgbot/methods"
	"tgbot/tgtype"
)

// BotAPIURL the URL bots use to communicate with telegram
const BotAPIURL = "https://api.telegram.org/bot"

// Bot stores info about a single bot
type Bot struct {
	OnMessage           func(tgtype.Message)
	OnMessageEdited     func(tgtype.Message)
	OnChannelPost       func(tgtype.Message)
	onChannelPostEdited func(tgtype.Message)
	APIKey              string
	running             bool
}

// GetBotURL - returns a telegram api call url only lacking method
func (bot Bot) GetBotURL() string {
	return fmt.Sprintf("%v%v/", string(BotAPIURL), bot.APIKey)
}

// RunBot starts the bot and runs it until bot gets stopped
func (bot *Bot) RunBot() {
	if bot.running {
		fmt.Println("Tried to launch an already running bot")
		return
	}
	bot.running = true
	for {
		method := methods.GetUpdates{
			Offset: 32,
			Limit:  5,
		}
		method.CallMethod(bot.GetBotURL())
	}
}
