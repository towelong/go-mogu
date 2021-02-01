package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"towelong/mogu/model"
)

const (
	url = "https://api.moguding.net:9000"
)

// MoGuService generate a serious of function's interfaces.
type MoGuService interface {
	MoGuLogin(account, password string) string
	GetPlanID(token string) string
	SignIn(token, planID string) bool
}

// moGuService is a empty struction, which include a serious of functions.
// mainly, in order to let it face to Object.
type moGuService struct {
}

// NewMoGuService init struction.
func NewMoGuService() MoGuService {
	return new(moGuService)
}

// MoGuLogin is a login logic of MoGu.
// account: When user register application,it is usually a phone number.
// password: Create by User
func (m moGuService) MoGuLogin(account, password string) string {
	body := map[string]string{
		"phone":     account,
		"password":  password,
		"loginType": "android",
		"uuid":      "",
	}
	client := &http.Client{}
	form, _ := json.Marshal(body)
	request, err := http.NewRequest(
		"POST",
		url+"/session/user/v1/login",
		bytes.NewReader(form),
	)
	if err == nil {
		request.Header.Add("accept-language", "zh-CN,zh;q=0.8")
		request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 9; zh-cn; ONEPLUS A6010 Build/PKQ1.180716.001) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Mobile Safari/533.1")
		request.Header.Add("content-type", "application/json; charset=UTF-8")
		request.Header.Add("cache-control", "no-cache")
		resp, error := client.Do(request)
		if error == nil {
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			var data model.DataModel
			json.Unmarshal(result, &data)
			if data.Code == 200 {
				return data.Data.Token
			}
		}
	}
	return ""
}

// getPlanID get task id
func (m moGuService) GetPlanID(token string) string {
	body := map[string]string{
		"paramsType": "student",
	}
	client := &http.Client{}
	form, _ := json.Marshal(body)
	request, err := http.NewRequest(
		"POST",
		url+"/practice/plan/v1/getPlanByStu",
		bytes.NewReader(form),
	)
	if err == nil {
		request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 9; zh-cn; ONEPLUS A6010 Build/PKQ1.180716.001) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Mobile Safari/533.1")
		request.Header.Add("content-type", "application/json; charset=UTF-8")
		request.Header.Add("Authorization", token)
		request.Header.Add("roleKey", "student")
		resp, error := client.Do(request)
		if error == nil {
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			var data model.PlanModel
			json.Unmarshal(result, &data)
			if data.Code == 200 {
				return data.Data[0].PlanID
			}
		}
	}
	return ""
}

// SignIn signIn Logic
func (m moGuService) SignIn(token, planID string) bool {
	types := "START"
	address := os.Getenv("ADDRESS")
	city := os.Getenv("CITY")
	province := os.Getenv("PROVINCE")
	longitude := os.Getenv("LONGITUDE")
	latitude := os.Getenv("LATITUDE")
	if address == "" && longitude == "" && city == "" {
		log.Fatal("failed to Load secret ")
	}
	utcHour := time.Now().UTC().Hour() + 8
	// I will go off work at 18:00
	if utcHour >= 12 && utcHour < 23 {
		// go off work sign
		types = "END"
	}
	body := &model.SignInModel{
		Device:         "Android",
		PlanID:         planID,
		Country:        "中国",
		Type:           types, // 默认打卡上班
		AttendanceType: "",
		State:          "NORMAL",
		Address:        address,
		Longitude:      longitude,
		Latitude:       latitude,
		City:           city,
		Province:       province,
	}
	client := &http.Client{}
	form, _ := json.Marshal(body)
	request, err := http.NewRequest(
		"POST",
		url+"/attendence/clock/v1/save",
		bytes.NewReader(form),
	)
	if err == nil {
		request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 9; zh-cn; ONEPLUS A6010 Build/PKQ1.180716.001) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Mobile Safari/533.1")
		request.Header.Add("content-type", "application/json; charset=UTF-8")
		request.Header.Add("Authorization", token)
		resp, error := client.Do(request)
		if error == nil {
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			var data map[string]interface{}
			json.Unmarshal(result, &data)
			if strings.EqualFold(data["msg"].(string), "success") {
				return true
			}
			fmt.Println(data["msg"])
		}
	}
	return false
}
