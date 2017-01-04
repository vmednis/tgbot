package methods

import "tgbot/tgtype"
import "encoding/json"
import "fmt"

// GetUpdates - Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.
type GetUpdates struct {
	Offset         int32    `json:"offset,omitempty"`
	Limit          int32    `json:"limit,omitempty"`
	Timeout        int32    `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_udpdates,omitempty"`
}

// GetName - returns the telegram calling name of the method
func (method GetUpdates) GetName() string {
	return "getUpdates"
}

// CallMethod - executes method returns update
func (method *GetUpdates) CallMethod(botURL string) tgtype.TGType {
	//Call the message under the hood
	result := callMethod(botURL, method).Result
	fmt.Println(string(result))

	//Decode the result
	updateSlice := make([]tgtype.Update, 0)
	json.Unmarshal([]byte(result), &updateSlice)

	var updates tgtype.Updates
	updates.Updates = updateSlice
	fmt.Println(updates.Updates[0].Message)
	return updates
}
