package viewmodels

type GetUsersListReq struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

type UserInfo struct {
	UserID      int64  `json:"user_id"`
	Username    string `json:"username"`
	Role        string `json:"role"`
	DeviceCount int64  `json:"device_count"`
	TotalPhones int64  `json:"total_phones"`
	CreatedAt   string `json:"created_at"`
}

type GetUsersListRsp struct {
	Users []UserInfo `json:"users"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
}

type GetDevicesListReq struct {
	Status string `json:"status"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
}

type DeviceInfo struct {
	DeviceID           string `json:"device_id"`
	UserID             *int64 `json:"user_id"`
	Username           string `json:"username"`
	TotalPhones        int64  `json:"total_phones"`
	PendingCount       int64  `json:"pending_count"`
	CalledCount        int64  `json:"called_count"`
	InterestedCount    int64  `json:"interested_count"`
	NotInterestedCount int64  `json:"not_interested_count"`
	CreatedAt          string `json:"created_at"`
}

type GetDevicesListRsp struct {
	Devices []DeviceInfo `json:"devices"`
	Total   int64        `json:"total"`
	Page    int          `json:"page"`
	Limit   int          `json:"limit"`
	Bound   int64        `json:"bound"`
	Unbound int64        `json:"unbound"`
}

type UnbindDeviceReq struct {
	DeviceID string `json:"device_id" binding:"required"`
}

type UnbindDeviceRsp struct {
	DeviceID string `json:"device_id"`
	Message  string `json:"message"`
}
