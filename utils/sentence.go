package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"towelong/mogu/model"
)

// RandomSentence 生成周报内容
func RandomSentence() (string, error) {
	rand.Seed(time.Now().UnixNano())
	var sentence model.SentenceModel
	resp, err := http.Get("https://raw.githubusercontent.com/ToWeLong/go-mogu/main/model/sentence.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &sentence)
	if len(sentence.Data) == 0 {
		return "", errors.New("数据总数为0")
	}
	r := rand.Intn(len(sentence.Data))
	str := sentence.Data[r].Text
	return str, nil
}
