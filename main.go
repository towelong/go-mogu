package main

import (
	"fmt"
	"towelong/mogu/service"
	"github.com/joho/godotenv"
)

func main() {
	// load enviroment secret key
	godotenv.Load()
	moguding := service.NewMoGuService()
	token := moguding.MoGuLogin("18779049477", "Fl1191430240.")
	planID := moguding.GetPlanID(token)
	isSuccess := moguding.SignIn(token, planID)
	if isSuccess {
		fmt.Println("打卡成功")
	}
}
