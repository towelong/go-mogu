package test

import (
	"fmt"
	"testing"
	"time"
	"towelong/mogu/utils"
)

func TestTime(t *testing.T) {
	types := "START"
	fmt.Println("local time:")
	fmt.Println(time.Now().Format("2006/1/2 15:04:05"))
	if time.Now().Hour() >= 12 {
		types = "END"
		fmt.Println("下班打卡啦~")
	}
	fmt.Println(types)
}

func TestWeek(t *testing.T) {
	utils.WeeklyPicker(time.Now())
	fmt.Println(utils.RandomSentence())
}

func TestRemoteTime(t *testing.T) {
	fmt.Println("local time:")
	fmt.Println(time.Now().Format("2006/1/2 15:04:05"))
	fmt.Println("UTC time:")
	fmt.Println(time.Now().UTC().Format("2006/1/2 15:04:05"))
	types := utils.TimePicker()
	fmt.Println(types)
}
