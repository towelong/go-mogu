package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	types := "START"
	fmt.Println(time.Now().Hour())
	if time.Now().Hour()>18{
		types = "END"
		fmt.Println("下班打卡啦~")
	}
	fmt.Println(types)
}
