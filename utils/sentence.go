package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"
	"towelong/mogu/model"
)

// RandomSentence 生成周报内容
func RandomSentence() (string, error) {
	rand.Seed(time.Now().UnixNano())
	var sentence model.SentenceModel
	abs, _ := os.Getwd()
	dir := filepath.Dir(abs)
	fmt.Println(dir)
	absPath := path.Join(dir, "/model/sentence.json")
	fmt.Println(absPath)
	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	jsonError := json.Unmarshal([]byte(file), &sentence)
	if jsonError != nil {
		return "", errors.New("JSON转换异常")
	}
	fmt.Println(len(sentence.Data))
	if len(sentence.Data) == 0 {
		return "", errors.New("数据总数为0")
	}
	r := rand.Intn(len(sentence.Data))
	str := sentence.Data[r].Text
	return str, nil
}
