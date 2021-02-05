package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"towelong/mogu/model"
)

// RandomSentence 生成周报内容
func RandomSentence() (string, error) {
	rand.Seed(time.Now().UnixNano())
	var sentence model.SentenceModel
	data, _ := ioutil.ReadFile("../model/sentence.json")
	json.Unmarshal([]byte(data), &sentence)
	fmt.Println(len(sentence.Data))
	if len(sentence.Data) == 0 {
		return "", errors.New("数据总数为0")
	}
	r := rand.Intn(len(sentence.Data))
	str := sentence.Data[r].Text
	return str, nil
}
