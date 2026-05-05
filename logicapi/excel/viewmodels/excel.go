package viewmodels

import "mime/multipart"

type UploadPhoneExcelReq struct {
	UserId string `json:"user_id"`
	// swagger:ignore
	File *multipart.FileHeader `form:"file"`
}

type UploadPhoneExcelRsp struct {
	List           []*PhoneStatus `json:"list"`
	EffectiveCount int            `json:"effective_count"` // 有效数量
	RepeatCount    int            `json:"repeat"`
	Total          int            `json:"int"` // 总共数量
}

type PhoneStatus struct {
	PhoneNumber int    `json:"phone_number"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
}
