package main

import "tgbot/bot"
import "tgbot/tgtype"
import "fmt"
import "tgbot/methods"

func main() {
	mednisBot := bot.Bot{APIKey: "320490619:AAGd1-Wtx-BMciXNxAV7gahNry8Wpzy8Guk"}

	mednisBot.OnMessage = func(msg *tgtype.Message) {
		fmt.Printf("Recieved message \"%v\" from user with name %v\n", msg.Text, msg.From.FirstName)
		fmt.Println(msg.Chat)
		method := methods.SendMessage{
			ChatID:           msg.Chat.ID,
			Text:             "STFU",
			ReplyToMessageID: msg.MessageID,
		}
		method.CallMethod(mednisBot.GetBotURL())
	}

	mednisBot.RunBot()
}
