package utils

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
	"towelong/mogu/model"
)

// RandomSentence 生成周报内容
func RandomSentence() string {
	rand.Seed(time.Now().UnixNano())
	var sentence model.SentenceModel
	data, _ := ioutil.ReadFile("../model/sentence.json")
	json.Unmarshal([]byte(data), &sentence)
	r := rand.Intn(len(sentence.Data))
	str := sentence.Data[r].Text
	return str
}
