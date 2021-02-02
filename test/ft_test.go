package test

import (
	"fmt"
	"testing"
	"towelong/mogu/service"
	"towelong/mogu/utils"

	"github.com/joho/godotenv"
)

func TestFT(t *testing.T) {
	godotenv.Load("../.env")
	types := utils.TimePicker()
	title, message := utils.EnumToMsg(types)
	b := service.SendMessage(title, message)
	if b {
		fmt.Println("Push Success!")
	}
	fmt.Println(b)
}
