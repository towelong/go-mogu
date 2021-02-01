package test

import (
	"fmt"
	"testing"
	"time"
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

func TestRemoteTime(t *testing.T) {
	types := "START"
	fmt.Println("local time:")
	fmt.Println(time.Now().Format("2006/1/2 15:04:05"))
	if time.Now().Hour()+8 >= 12 {
		types = "END"
		fmt.Println("下班打卡啦~")
	}
	fmt.Println(types)
}
