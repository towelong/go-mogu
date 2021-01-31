package model

// DataModel is used to reflect data
type DataModel struct {
	Code int  `json:"code"`
	Data Data `json:"data"`
}

// Data is used reflect data to dataDodel
type Data struct {
	NiKeName string `json:"nikeName"`
	UserID   string `json:"userId"`
	Token    string `json:"token"`
}

// PlanModel return Object
type PlanModel struct {
	Code int    `json:"code"`
	Data []Plan `json:"data"`
}

// Plan Object
type Plan struct {
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
	AttendanceType string `json:"attendanceType"`
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
