package service

import (
	"fmt"
	"net/http"
	"os"
)

const (
	serverChanAPI = "https://sc.ftqq.com/"
)

// SendMessage we can use Server Chan to send message.
// text: message title.
// desp: message content.
func SendMessage(text, desp string) bool {
	secretKey := os.Getenv("KEY")
	urlStr := fmt.Sprintf(serverChanAPI+"%v.send?text=%v&desp=%v", secretKey, text, desp)
	fmt.Println(urlStr)
	if resp, err := http.Get(urlStr); err == nil {
		defer resp.Body.Close()
		return true
	}
	return false
}
