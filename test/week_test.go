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
	token := moguding.MoGuLogin(os.Getenv("ACCOUNT"), os.Getenv("PASSWORD"))
	planID := moguding.GetPlanID(token)
	fmt.Println(moguding.WeeklyDiary(token, planID))
}
