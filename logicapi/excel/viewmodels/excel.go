package viewmodels

import "mime/multipart"

// ===== 接口：上传Excel接口 =====

type UploadPhoneExcelReq struct {
	UserId   string `form:"user_id" json:"user_id"`
	DeviceID string `form:"device_id" json:"device_id"`
	// swagger:ignore
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadPhoneExcelRsp struct {
	PhoneList      []PhoneInfo `json:"phone_list"`      // 电话号码列表
	EffectiveCount int         `json:"effective_count"` // 有效数量
	RepeatCount    int         `json:"repeat_count"`    // 重复数量
	Total          int         `json:"total"`           // 总共数量
}

type PhoneInfo struct {
	ID          int64  `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
}

// ===== 接口：读取电话号码列表接口 =====
type GetPhonesReq struct {
	DeviceID string `json:"device_id"` // 设备ID
}

type GetPhonesRsp struct {
	PhoneList []PhoneInfo `json:"phone_list"` // 电话号码列表
	Total     int         `json:"total"`      // 总数
}

type GetDevicesReq struct {
}

type GetDevicesRsp struct {
	DeviceList []string `json:"device_list"`
	Total      int      `json:"total"`
}

// ===== 接口：软删除电话号码 =====
type DeletePhoneReq struct {
	ID int64 `json:"id" binding:"required"` // 电话号码记录ID
}

type DeletePhoneRsp struct {
	ID int64 `json:"id"` // 已删除的记录ID
}

// ===== 接口：编辑电话号码备注/状态 =====
type UpdatePhoneReq struct {
	ID     int64   `json:"id" binding:"required"` // 电话号码记录ID
	Remark *string `json:"remark"`                // 备注（可选）
	Status *int    `json:"status"`                // 状态（可选）
}

type UpdatePhoneRsp struct {
	PhoneInfo
}

type RegisterDeviceReq struct {
	DeviceID string `json:"device_id" binding:"required"` // 设备ID
}

type RegisterDeviceRsp struct {
	DeviceID string `json:"device_id"`
	Created  bool   `json:"created"` // 是否新创建设备记录
}

type UpdateClientPhoneStatusReq struct {
	DeviceID    string `json:"device_id,omitempty"`             // 设备ID，可选，若全局号码不唯一则需传入
	PhoneNumber string `json:"phone_number" binding:"required"` // 电话号码
	Status      int    `json:"status" binding:"required"`       // 新状态，0-待拨打，1-已拨打，2-有意向，3-无意向
}

type UpdateClientPhoneStatusRsp struct {
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}

// ===== 接口：C端返回电话号码数组 ======

type GetClientPhoneReq struct {
	DeviceID string `json:"device_id"`
}

type ClientPhoneInfo struct {
	PhoneNumber string `json:"phone_number"`
	CompanyName string `json:"company_name"`
}

type GetClentPhoneRsp struct {
	Phones []ClientPhoneInfo `json:"phones"`
}

// ===== 接口：绑定设备到用户 =====

type BindDeviceReq struct {
	DeviceID string `json:"device_id" binding:"required"`
	UserID   int64  `json:"user_id" binding:"required"`
}

type BindDeviceRsp struct {
	DeviceID string `json:"device_id"`
	UserID   int64  `json:"user_id"`
}

// ===== 接口：概览统计 =====

type OverviewRsp struct {
	TotalDevices   int         `json:"total_devices"`
	TotalPhones    int64       `json:"total_phones"`
	TotalUsers     int64       `json:"total_users"`
	UnboundDevices int64       `json:"unbound_devices"`
	PhonesByStatus map[int]int `json:"phones_by_status"`
}

// ===== 接口：设备统计列表 =====

type DeviceStats struct {
	DeviceID           string `json:"device_id"`
	UserName           string `json:"user_name"`
	UserID             *int64 `json:"user_id"`
	TotalPhones        int64  `json:"total_phones"`
	PendingCount       int64  `json:"pending_count"`
	CalledCount        int64  `json:"called_count"`
	InterestedCount    int64  `json:"interested_count"`
	NotInterestedCount int64  `json:"not_interested_count"`
}

type GetDevicesStatsRsp struct {
	Devices []DeviceStats `json:"devices"`
	Total   int           `json:"total"`
}
