package main

import (
	"fmt"
	"log"
	"os"
	"towelong/mogu/service"
	"towelong/mogu/utils"

	"github.com/joho/godotenv"
)

func main() {
	// load enviroment secret key
	godotenv.Load()
	moguding := service.NewMoGuService()
	token, userId := moguding.MoGuLogin(os.Getenv("ACCOUNT"), os.Getenv("PASSWORD"))
	planID := moguding.GetPlanID(token, userId)
	isSuccess, types := moguding.SignIn(token, planID, userId)
	title, message := utils.EnumToMsg(types)
	if !isSuccess {
		service.SendMessage(title, message)
		log.Fatal(title)
	}
	if service.SendMessage(title, message) {
		fmt.Println("打卡成功")
	}

	// 写周记
	isTrue, weekType := moguding.WeeklyDiary(token, planID)
	headline, msg := utils.EnumToMsg(weekType)
	if !isTrue && weekType == utils.NOWEEK {
		fmt.Println(msg)
	} else {
		if service.SendMessage(headline, msg) {
			fmt.Println("打卡成功")
		}
	}
}
