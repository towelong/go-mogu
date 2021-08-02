package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"towelong/mogu/model"
	"towelong/mogu/utils"
)

const (
	url  = "https://api.moguding.net:9000"
	salt = "3478cbbc33f84bd00d75d7dfa69e0daa"
)

// MoGuService generate a serious of function's interfaces.
type MoGuService interface {
	MoGuLogin(account, password string) (string, string)
	GetPlanID(token, userId string) string
	SignIn(token, planID, userId string) (bool, string)
	WeeklyDiary(token, planID string) (bool, string)
}

// moGuService is a empty struction, which include a serious of functions.
// mainly, in order to let it face to Object.
type moGuService struct {
}

// NewMoGuService init struction.
func NewMoGuService() MoGuService {
	return &moGuService{}
}

// MoGuLogin is a login logic of MoGu.
// account: When user register application,it is usually a phone number.
// password: Create by User
func (m moGuService) MoGuLogin(account, password string) (token, userId string) {
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
		request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 4.4.2; en-us; Android SDK built for x86 Build/KK) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")
		request.Header.Add("content-type", "application/json; charset=UTF-8")
		request.Header.Add("cache-control", "no-cache")
		resp, err := client.Do(request)
		if err == nil {
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			var data model.DataModel
			_ = json.Unmarshal(result, &data)
			if data.Code == 200 {
				return data.Data.Token, data.Data.UserID
			}
		}
	}
	return "", ""
}

// GetPlanID getPlanID get task id
func (m moGuService) GetPlanID(token, userId string) string {
	body := map[string]int{
		"state": 1,
	}
	client := &http.Client{}
	form, _ := json.Marshal(body)
	request, err := http.NewRequest(
		"POST",
		url+"/practice/plan/v3/getPlanByStu",
		bytes.NewReader(form),
	)
	if err == nil {
		// 计算获取实习计划的sign
		// str = userId + student + salt
		sign := utils.CreateSign(userId + "student" + salt)
		request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 9; zh-cn; ONEPLUS A6010 Build/PKQ1.180716.001) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Mobile Safari/533.1")
		request.Header.Add("content-type", "application/json; charset=UTF-8")
		request.Header.Add("authorization", token)
		request.Header.Add("rolekey", "student")
		request.Header.Add("sign", sign)
		resp, err := client.Do(request)
		if err == nil {
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			var data model.PlanModel
			_ = json.Unmarshal(result, &data)
			if data.Code == 200 {
				return data.Data[0].PlanID
			}
		}
	}
	return ""
}

// SignIn signIn Logic
func (m moGuService) SignIn(token, planID, userId string) (bool, string) {
	address := os.Getenv("ADDRESS")
	city := os.Getenv("CITY")
	province := os.Getenv("PROVINCE")
	longitude := os.Getenv("LONGITUDE")
	latitude := os.Getenv("LATITUDE")
	if address == "" && longitude == "" && city == "" {
		log.Fatal("failed to Load secret ")
	}
	// 自动计算 上午 or 下午
	// 上午为 上班打卡;下午为 下班打卡
	types := utils.TimePicker()
	body := &model.SignInModel{
		Device:      "Android",
		PlanID:      planID,
		Country:     "中国",
		Type:        types, // 默认打卡上班
		Description: "",
		State:       "NORMAL",
		Address:     address,
		Longitude:   longitude,
		Latitude:    latitude,
		City:        city,
		Province:    province,
	}
	client := &http.Client{}
	form, _ := json.Marshal(body)
	request, err := http.NewRequest(
		"POST",
		url+"/attendence/clock/v2/save",
		bytes.NewReader(form),
	)
	if err == nil {
		str := body.Device + types + planID + userId + body.Address + salt
		sign := utils.CreateSign(str)
		request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 9; zh-cn; ONEPLUS A6010 Build/PKQ1.180716.001) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Mobile Safari/533.1")
		request.Header.Add("content-type", "application/json; charset=UTF-8")
		request.Header.Add("Authorization", token)
		request.Header.Add("sign", sign)
		resp, err := client.Do(request)
		if err == nil {
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			var data map[string]interface{}
			_ = json.Unmarshal(result, &data)
			if data["code"].(float64) == 200 {
				return true, types
			}
			fmt.Println(data["msg"])
		}
	}
	return false, utils.ERROR
}

// WeeklyDiary it will be automatic writing weekly diary.
func (m moGuService) WeeklyDiary(token, planID string) (bool, string) {
	if !(time.Now().UTC().Weekday() == time.Saturday && utils.TimePicker() == utils.END) {
		return false, utils.NOWEEK
	}
	sentence, randomErr := utils.RandomSentence()
	fmt.Println(sentence)
	if randomErr != nil {
		log.Fatal(randomErr)
	}
	currentWeek, startTime, endTime := utils.WeeklyPicker(time.Now())
	body := &model.WeekWriterModel{
		AttachmentList: []string{},
		Attachments:    "",
		PlanID:         planID,
		ReportType:     "week",
		Title:          fmt.Sprintf("第%v周周报", currentWeek),
		Content:        sentence,
		Weeks:          fmt.Sprintf("第%v周", currentWeek),
		StartTime:      startTime,
		EndTime:        endTime,
	}
	client := &http.Client{}
	form, _ := json.Marshal(body)
	request, err := http.NewRequest(
		"POST",
		url+"/practice/paper/v1/save",
		bytes.NewReader(form),
	)
	if err == nil {
		request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 9; zh-cn; ONEPLUS A6010 Build/PKQ1.180716.001) AppleWebKit/533.1 (KHTML, like Gecko) Version/5.0 Mobile Safari/533.1")
		request.Header.Add("content-type", "application/json; charset=UTF-8")
		request.Header.Add("Authorization", token)
		resp, err := client.Do(request)
		if err == nil {
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			var data map[string]interface{}
			_ = json.Unmarshal(result, &data)
			if data["code"].(float64) == 200 {
				return true, utils.WEEK
			}
			if data["code"].(float64) == 500 {
				fmt.Println(data["msg"])
				return false, utils.NOWEEK
			}
		}
	}
	return false, utils.ERROR
}
