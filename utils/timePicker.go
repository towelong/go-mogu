package utils

import "time"

// TimePicker 转成UTC时间进行判断是否上班和下班
func TimePicker() string {
	utcHour := time.Now().UTC().Hour() + 8
	// I will go off work at 18:00
	if utcHour >= 12 && utcHour <= 23 {
		// go off work sign
		return END
	}
	return START
}
