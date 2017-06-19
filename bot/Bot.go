package bot

import (
	"fmt"

	"github.com/vmednis/tgbot/commandprocessor"
	"github.com/vmednis/tgbot/methods"
	"github.com/vmednis/tgbot/tgtype"
)

// BotAPIURL the URL bots use to communicate with telegram
const BotAPIURL = "https://api.telegram.org/bot"

// Bot stores info about a single bot
type Bot struct {
	OnMessage           func(*tgtype.Message)
	OnMessageEdited     func(*tgtype.Message)
	OnChannelPost       func(*tgtype.Message)
	OnChannelPostEdited func(*tgtype.Message)
	CommandProcessor    commandprocessor.CommandPorcessor
	APIKey              string
	running             bool
	updateOffset        int64
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
		//Request the new updates
		method := methods.GetUpdates{
			Offset:  bot.updateOffset,
			Limit:   5,
			Timeout: 60,
		}
		var updates tgtype.Updates
		updates = method.CallMethod(bot.GetBotURL()).(tgtype.Updates)

		//Parse the new updates
		for _, u := range updates.Updates {
			switch {
			case u.Message != nil:
				if !bot.CommandProcessor.ExecuteCommand(*u.Message) { //If the message isn't a command used up by the command processor
					if bot.OnMessage != nil {
						bot.OnMessage(u.Message)
					}
				}
			case u.EditedMessage != nil:
				if bot.OnMessageEdited != nil {
					bot.OnMessageEdited(u.EditedMessage)
				}
			case u.ChannelPost != nil:
				if bot.OnChannelPost != nil {
					bot.OnChannelPost(u.ChannelPost)
				}
			case u.EditedChannelPost != nil:
				if bot.OnChannelPostEdited != nil {
					bot.OnChannelPostEdited(u.EditedChannelPost)
				}
			default:
				fmt.Println("Receieved an unrecognized update!")
			}

			if u.UpdateID >= bot.updateOffset {
				bot.updateOffset = u.UpdateID + 1
			}
		}
	}
}
