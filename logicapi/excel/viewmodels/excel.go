package viewmodels

import "mime/multipart"

type UploadPhoneExcelReq struct {
	UserId string `json:"user_id"`
	// swagger:ignore
	File *multipart.FileHeader `form:"file"`
}

type UploadPhoneExcelRsp struct {
	PhoneList      []PhoneInfo `json:"phone_list"`
	EffectiveCount int         `json:"effective_count"` // 有效数量
	RepeatCount    int         `json:"repeat"`
	Total          int         `json:"int"` // 总共数量
}

type PhoneInfo struct {
	PhoneNumber string `json:"phone_number"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
}
