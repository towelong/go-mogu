package utils

import "strings"

const (
	// START 上班
	START = "START"
	// END 下班
	END = "END"
	// ERROR 运行错误
	ERROR = "ERROR"
)

var statusMsg = map[string]string{
	START: "上班打卡成功~ 上班签到成功",
	END:   "下班打卡成功~ 下班签到成功",
	ERROR: "系统崩溃了~ o(╥﹏╥)o",
}

// EnumToMsg 枚举转换成具有意义的字符串
func EnumToMsg(enumValue string) (string, string) {
	str, ok := statusMsg[enumValue]
	if ok {
		slice := strings.Fields(str)
		return slice[0], slice[1]
	}
	errMsg := strings.Fields(statusMsg[ERROR])
	return errMsg[0], errMsg[1]
}
