package main

import (
	"fmt"
	"os"
	"towelong/mogu/service"

	"github.com/joho/godotenv"
)

func main() {
	// load enviroment secret key
	godotenv.Load()
	moguding := service.NewMoGuService()
	token := moguding.MoGuLogin(os.Getenv("ACCOUNT"), os.Getenv("PASSWORD"))
	planID := moguding.GetPlanID(token)
	isSuccess := moguding.SignIn(token, planID)
	if isSuccess {
		fmt.Println("打卡成功")
	}
}
