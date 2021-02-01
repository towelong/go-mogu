package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// local test
func TestLocalEnv(t *testing.T) {
	err := godotenv.Load("../.env")
	if err == nil {
		fmt.Println(os.Getenv("ACCOUNT"))
		fmt.Println(os.Getenv("PASSWORD"))
		fmt.Println(os.Getenv("ADDRESS"))
		fmt.Println(os.Getenv("KEY"))
	} else {
		fmt.Println("环境变量读取失败")
	}
}

// remote test
func TestRemoteEnv(t *testing.T) {
	address := os.Getenv("ADDRESS")
	city := os.Getenv("CITY")
	key := os.Getenv("KEY")
	if address == "" && city == "" && key == "" {
		fmt.Println("failed to Load secret ")
	} else {
		fmt.Println("Load secret success")
	}
}
