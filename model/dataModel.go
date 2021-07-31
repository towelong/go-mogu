package model

// DataModel is used to reflect data
type DataModel struct {
	Code int  `json:"code"`
	Data data `json:"data"`
}

// Data is used reflect data to dataDodel
type data struct {
	NiKeName string `json:"nikeName"`
	UserID   string `json:"userId"`
	Token    string `json:"token"`
}

// PlanModel return Object
type PlanModel struct {
	Code int    `json:"code"`
	Data []plan `json:"data"`
}

// Plan Object
type plan struct {
	PlanID   string `json:"planId"`
	PlanName string `json:"planName"`
}

// SignInModel signIn form data
type SignInModel struct {
	// sign device
	Device string `json:"device"`
	// planId
	PlanID string `json:"planId"`
	// default: China
	Country string `json:"country"`
	// default: NORMAL
	State string `json:"state"`
	// default:""
	Description string `json:"description"`
	// address
	Address string `json:"address"`
	// "go to work"：START or "go off work":END
	Type string `json:"type"`
	// 经度 参考：https://lbs.amap.com/console/show/picker
	Longitude string `json:"longitude"`
	// 纬度 参考：https://lbs.amap.com/console/show/picker
	Latitude string `json:"latitude"`
	// 区
	City string `json:"city"`
	// 市
	Province string `json:"province"`
}

// WeekWriterModel Weekly Diary form
type WeekWriterModel struct {
	AttachmentList []string `json:"attachmentList"`
	Attachments    string   `json:"attachments"`
	// 周报内容
	Content string `json:"content"`
	PlanID  string `json:"planId"`
	// 周报类型
	ReportType string `json:"reportType"`
	// 周报标题
	Title string `json:"title"`
	// 第 X 周
	Weeks string `json:"weeks"`
	// 当前周开始时间
	StartTime string `json:"startTime"`
	// 当前周结束时间
	EndTime string `json:"endTime"`
}

// SentenceModel init data
type SentenceModel struct {
	Code int    `json:"code"`
	Data []text `json:"data"`
}

type text struct {
	Text string `json:"text"`
}
