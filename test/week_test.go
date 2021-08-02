package test

import (
	"fmt"
	"os"
	"testing"
	"towelong/mogu/service"

	"github.com/joho/godotenv"
)

func TestWeekDiary(t *testing.T) {
	godotenv.Load("../.env")
	moguding := service.NewMoGuService()
	token, userId := moguding.MoGuLogin(os.Getenv("ACCOUNT"), os.Getenv("PASSWORD"))
	planID := moguding.GetPlanID(token, userId)
	fmt.Println(moguding.WeeklyDiary(token, planID))
}
