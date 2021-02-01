package test

import (
	"fmt"
	"testing"
	"time"
	"towelong/mogu/service"

	"github.com/joho/godotenv"
)

func TestFT(t *testing.T) {
	godotenv.Load("../.env")
	times := time.Now().Format("2006年1月2日15:04:05")
	message := fmt.Sprintf("打卡时间为%v", times)
	service.SendMessage("上班打卡成功提醒", message)
}
