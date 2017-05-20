package methods

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/vmednis/tgbot/tgtype"
)

// Method a telegram method
type Method interface {
	GetName() string
	CallMethod(botURL string) tgtype.TGType
}

// MethodResponse - struct returned by all the methods
type MethodResponse struct {
	OK          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"descriptopm,omitempty"`
}

func callMethod(botURL string, method Method) MethodResponse {
	var response MethodResponse

	//Encode the request body
	bodyBuf, _ := json.Marshal(method)
	requestBody := string(bodyBuf)

	//Create the request url
	url := fmt.Sprintf("%v%v", botURL, method.GetName())

	//Call the method
	cl := http.Client{
		Timeout: time.Duration(61 * time.Second),
	}
	resp, err := cl.Post(url, "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		fmt.Println("Error getting update: ", err)
	} else {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&response)
		if err != nil {
			fmt.Println("Error decoding json in callMethod()")
		}
	}

	return response
}
