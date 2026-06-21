package datamodels

import "time"

type UserDevice struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DeviceID  string    `gorm:"column:device_id;type:varchar(128);not null;uniqueIndex:uk_user_device" json:"device_id"`
	UserID    *int64    `gorm:"column:user_id;index" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (UserDevice) TableName() string {
	return "user_device"
}
