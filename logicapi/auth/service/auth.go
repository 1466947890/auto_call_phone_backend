package service

import (
	"auto_call_phone/data/datamodels"
	"auto_call_phone/data/repo"
	vm "auto_call_phone/logicapi/auth/viewmodels"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, req *vm.RegisterReq) (*vm.RegisterRsp, error) {
	if len(req.Username) < 3 || len(req.Username) > 64 {
		return nil, errors.New("username must be 3-64 characters")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &datamodels.User{
		Username:     req.Username,
		PasswordHash: string(hash),
	}
	if err := repo.CreateUser(ctx, user); err != nil {
		return nil, errors.New("username already exists")
	}
	return &vm.RegisterRsp{ID: user.ID, Username: user.Username}, nil
}

func (s *AuthService) Login(ctx context.Context, req *vm.LoginReq) (int64, *vm.LoginRsp, error) {
	user, err := repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil, errors.New("invalid username or password")
		}
		return 0, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return 0, nil, errors.New("invalid username or password")
	}

	return user.ID, &vm.LoginRsp{Username: user.Username}, nil
}

func (s *AuthService) GetUserDevices(ctx context.Context, userID int64) (*vm.GetUserDevicesRsp, error) {
	rows, err := repo.GetUserDevices(ctx, userID)
	if err != nil {
		return nil, err
	}
	devices := make([]string, 0, len(rows))
	for _, r := range rows {
		devices = append(devices, r.DeviceID)
	}
	return &vm.GetUserDevicesRsp{Devices: devices, Total: len(devices)}, nil
}
