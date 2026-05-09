package datamodels

import "time"

type DevicePhone struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(128);not null;uniqueIndex:uk_phone_device,priority:1" json:"phone_number"`
	DeviceID    string    `gorm:"column:device_id;type:varchar(128);not null;uniqueIndex:uk_phone_device,priority:2" json:"device_id"`
	Status      int8      `gorm:"column:status;type:tinyint(4);not null;default:0" json:"status"`
	Remarks     string    `gorm:"column:remarks;type:varchar(256);not null;default:''" json:"remarks"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (DevicePhone) TableName() string {
	return "device_phone"
}
