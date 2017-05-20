package methods

import (
	"encoding/json"

	"github.com/vmednis/tgbot/tgtype"
)

// GetMe - A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
type GetMe struct {
}

// GetName - returns the telegram calling name of the method
func (method GetMe) GetName() string {
	return "getMe"
}

// CallMethod - executes method returns update
// Returns User
func (method *GetMe) CallMethod(botURL string) tgtype.TGType {
	//Call the message under the hood
	result := callMethod(botURL, method).Result

	//Decode the result
	var user tgtype.User
	json.Unmarshal([]byte(result), &user)

	return user
}
