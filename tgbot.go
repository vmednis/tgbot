package main

import "tgbot/bot"
import "tgbot/tgtype"
import "fmt"
import "tgbot/methods"

//Very basic, barebones example bot.
func main() {
    mednisBot := bot.Bot{APIKey: "APIkey:here"}

	mednisBot.OnMessage = func(msg *tgtype.Message) {
		fmt.Printf("Recieved message \"%v\" from user with name %v\n", msg.Text, msg.From.FirstName)
		fmt.Println(msg.Chat)
		method := methods.SendMessage{
			ChatID:           msg.Chat.ID,
			Text:             "Hello world!",
			ReplyToMessageID: msg.MessageID,
		}
		method.CallMethod(mednisBot.GetBotURL())
	}

	mednisBot.RunBot()
}
