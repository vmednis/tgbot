package methods

import "encoding/json"
import "tgbot/tgtype"

// SendMessage - A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
type SendMessage struct {
	ChatID                interface{} `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageID      int32       `json:"reply_to_message_id,omitempty"`
}

// GetName - returns the telegram calling name of the method
func (method SendMessage) GetName() string {
	return "sendMessage"
}

// CallMethod - executes method returns update
// Returns Message
func (method *SendMessage) CallMethod(botURL string) tgtype.TGType {
	//Call the message under the hood
	result := callMethod(botURL, method).Result

	//Decode the result
	var user tgtype.Message
	json.Unmarshal([]byte(result), &user)

	return user
}
