package datamodels

import "time"

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username     string    `gorm:"column:username;type:varchar(64);not null;uniqueIndex:uk_username" json:"username"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(256);not null" json:"-"`
	Role         string    `gorm:"column:role;type:varchar(16);not null;default:user" json:"role"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
