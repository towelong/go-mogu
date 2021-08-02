package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	serverChanAPI = "https://sctapi.ftqq.com/"
)

// SendMessage we can use Server Chan to send message.
// text: message title.
// desp: message content.
func SendMessage(text, desp string) bool {
	secretKey := os.Getenv("KEY")
	urlStr := fmt.Sprintf(serverChanAPI+"%v.send?text=%v&desp=%v", secretKey, text, desp)
	if resp, err := http.Get(urlStr); err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		res := make(map[string]interface{})
		json.Unmarshal(body, &res)
		if res["code"].(float64) == 0 {
			return true
		}
	}
	return false
}
