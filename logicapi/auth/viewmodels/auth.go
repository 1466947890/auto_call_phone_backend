package viewmodels

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRsp struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type GetUserDevicesRsp struct {
	Devices []string `json:"devices"`
	Total   int      `json:"total"`
}
