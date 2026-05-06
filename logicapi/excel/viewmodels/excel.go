package viewmodels

import "mime/multipart"

// ===== 接口：上传Excel接口 =====

type UploadPhoneExcelReq struct {
	UserId string `json:"user_id"`
	// swagger:ignore
	File *multipart.FileHeader `form:"file"`
}

type UploadPhoneExcelRsp struct {
	PhoneList      []PhoneInfo `json:"phone_list"`      // 电话号码列表
	EffectiveCount int         `json:"effective_count"` // 有效数量
	RepeatCount    int         `json:"repeat"`          // 重复数量
	Total          int         `json:"int"`             // 总共数量
}

type PhoneInfo struct {
	PhoneNumber string `json:"phone_number"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
}

// ===== 接口：读取电话号码列表接口 =====
type GetPhonesReq struct {
	DeviceID string `json:"device_id"` // 设备ID
}

type GetPhonesRsp struct {
	PhoneList []string `json:"phone_list"` // 电话号码列表
}
