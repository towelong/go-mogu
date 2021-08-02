package utils

import (
	"fmt"
	"time"
)

var (
	// WeekDay range[0-6]
	WeekDay = map[string]int{
		time.Monday.String():    1,
		time.Tuesday.String():   2,
		time.Wednesday.String(): 3,
		time.Thursday.String():  4,
		time.Friday.String():    5,
		time.Saturday.String():  6,
		time.Sunday.String():    7,
	}
)

// WeeklyPicker 计算当前为第几周，该周的第一天，该周的最后一天
// return currentWeek, mondayTime, sundayTime
func WeeklyPicker(end time.Time) (int, string, string) {
	// 周记开始时间
	startWeekTime := "2020-11-30"
	start, _ := time.Parse("2006-01-02", startWeekTime)
	currentWeek := int(end.Sub(start).Hours())/24/7 + 1
	weekDay := WeekDay[end.Weekday().String()]
	weekDay2 := WeekDay[end.Weekday().String()]
	fmt.Printf("当前为第%v周的第%v天\n", currentWeek, weekDay)
	// 计算离星期天还有几天
	last := 0
	for {
		if weekDay == 7 {
			break
		}
		weekDay++
		last++
	}
	// 计算离星期一还有几天
	first := 0
	for {
		if weekDay2 == 1 {
			break
		}
		weekDay2--
		first--
	}
	firstDay, _ := time.ParseDuration(fmt.Sprintf("%vh", first*24))
	lastDay, _ := time.ParseDuration(fmt.Sprintf("%vh", last*24))

	mondayTime := end.Add(firstDay).Format("2006-01-02") + " 00:00:00"
	sundayTime := end.Add(lastDay).Format("2006-01-02") + " 23:59:59"
	fmt.Printf("第%v周的第一天为%v\n第七天为%v\n", currentWeek, mondayTime, sundayTime)
	return currentWeek, mondayTime, sundayTime
}
