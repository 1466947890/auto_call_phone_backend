package service

import (
	"auto_call_phone/common/base"
	"auto_call_phone/data/repo"
	"auto_call_phone/logicapi/user/viewmodels"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (p UserService) GetUsersList(c context.Context, commonParams *base.CommonParams, req *viewmodels.GetUsersListReq) (*viewmodels.GetUsersListRsp, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 || req.Limit > 100 {
		req.Limit = 20
	}
	users, total, err := repo.GetUsersList(c, req.Keyword, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	userInfos := make([]viewmodels.UserInfo, 0, len(users))
	for _, user := range users {
		deviceCount, _ := repo.CountUserDevices(c, user.ID)
		totalPhones, _ := repo.CountUserPhones(c, user.ID)
		userInfos = append(userInfos, viewmodels.UserInfo{
			UserID:      user.ID,
			Username:    user.Username,
			Role:        user.Role,
			DeviceCount: deviceCount,
			TotalPhones: totalPhones,
			CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &viewmodels.GetUsersListRsp{
		Users: userInfos,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

func (p UserService) GetDevicesList(c context.Context, commonParams *base.CommonParams, req *viewmodels.GetDevicesListReq) (*viewmodels.GetDevicesListRsp, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 || req.Limit > 100 {
		req.Limit = 20
	}
	devices, total, boundCount, unboundCount, err := repo.GetDevicesList(c, req.Status, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	deviceInfos := make([]viewmodels.DeviceInfo, 0, len(devices))
	for _, d := range devices {
		username := ""
		if d.UserName != nil {
			username = *d.UserName
		}
		deviceInfos = append(deviceInfos, viewmodels.DeviceInfo{
			DeviceID:           d.DeviceID,
			UserID:             d.UserID,
			Username:           username,
			TotalPhones:        d.TotalPhones,
			PendingCount:       d.PendingCount,
			CalledCount:        d.CalledCount,
			InterestedCount:    d.InterestedCount,
			NotInterestedCount: d.NotInterestedCount,
			CreatedAt:          d.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &viewmodels.GetDevicesListRsp{
		Devices: deviceInfos,
		Total:   total,
		Page:    req.Page,
		Limit:   req.Limit,
		Bound:   boundCount,
		Unbound: unboundCount,
	}, nil
}

func (p UserService) UnbindDevice(c context.Context, commonParams *base.CommonParams, req *viewmodels.UnbindDeviceReq) (*viewmodels.UnbindDeviceRsp, error) {
	if req.DeviceID == "" {
		return nil, errors.New("device_id is required")
	}
	if err := repo.UnbindDevice(c, req.DeviceID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("device not found")
		}
		return nil, err
	}
	return &viewmodels.UnbindDeviceRsp{
		DeviceID: req.DeviceID,
		Message:  "Device unbound successfully",
	}, nil
}
