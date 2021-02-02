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
	token := moguding.MoGuLogin(os.Getenv("ACCOUNT"), os.Getenv("PASSWORD"))
	planID := moguding.GetPlanID(token)
	isSuccess, types := moguding.SignIn(token, planID)
	title, message := utils.EnumToMsg(types)
	if !isSuccess {
		service.SendMessage(title, message)
		log.Fatal(title)
	}
	service.SendMessage(title, message)
	fmt.Println("打卡成功")
}
