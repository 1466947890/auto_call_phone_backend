package repo

import (
	"auto_call_phone/data/datamodels"
	"context"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitGrom() error {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		host := getEnv("MYSQL_HOST", "192.168.31.77")
		port := getEnv("MYSQL_PORT", "3306")
		user := getEnv("MYSQL_USER", "auto_call_phone")
		pass := getEnv("MYSQL_PASSWORD", "aiHakWpFm3HdBFyz")
		dbname := getEnv("MYSQL_DATABASE", "auto_call_phone")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, dbname)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	if err := db.AutoMigrate(&datamodels.User{}, &datamodels.UserDevice{}, &datamodels.DevicePhone{}); err != nil {
		return err
	}
	return nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// GetPendingPhonesByDeviceID 获取指定设备ID的待处理电话号码列表
func GetPendingPhonesByDeviceID(ctx context.Context, deviceID string) ([]datamodels.DevicePhone, error) {
	var rows []datamodels.DevicePhone
	if err := DB.WithContext(ctx).
		Where("device_id = ? AND status = ?", deviceID, datamodels.PhoneStatusPending).
		Order("id ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

// CreateUser 创建用户
func CreateUser(ctx context.Context, user *datamodels.User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(ctx context.Context, username string) (*datamodels.User, error) {
	var user datamodels.User
	if err := DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserDevices 获取用户关联的设备ID列表
func GetUserDevices(ctx context.Context, userID int64) ([]datamodels.UserDevice, error) {
	var rows []datamodels.UserDevice
	if err := DB.WithContext(ctx).Where("user_id = ?", userID).Order("id ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

// BindDeviceToUser 绑定设备到用户
func BindDeviceToUser(ctx context.Context, deviceID string, userID int64) error {
	result := DB.WithContext(ctx).Model(&datamodels.UserDevice{}).
		Where("device_id = ?", deviceID).
		Update("user_id", userID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// CountDevices 设备总数
func CountDevices(ctx context.Context) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&datamodels.UserDevice{}).Count(&count).Error
	return count, err
}

// CountUsers 用户总数
func CountUsers(ctx context.Context) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&datamodels.User{}).Count(&count).Error
	return count, err
}

// CountPhones 电话号码总数（不含软删除）
func CountPhones(ctx context.Context) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&datamodels.DevicePhone{}).Count(&count).Error
	return count, err
}

// CountUnboundDevices 未绑定用户的设备数
func CountUnboundDevices(ctx context.Context) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&datamodels.UserDevice{}).Where("user_id IS NULL").Count(&count).Error
	return count, err
}

// CountPhonesByStatus 按状态统计电话号码数量
func CountPhonesByStatus(ctx context.Context) (map[int]int64, error) {
	type row struct {
		Status int   `gorm:"column:status"`
		Count  int64 `gorm:"column:count"`
	}
	var rows []row
	err := DB.WithContext(ctx).Model(&datamodels.DevicePhone{}).
		Select("status, count(*) as count").
		Group("status").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	result := make(map[int]int64)
	for _, r := range rows {
		result[r.Status] = r.Count
	}
	return result, nil
}

// GetUserByID 根据ID获取用户
func GetUserByID(ctx context.Context, userID int64) (*datamodels.User, error) {
	var user datamodels.User
	if err := DB.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeviceStatsRow 设备统计查询结果行
type DeviceStatsRow struct {
	DeviceID           string  `gorm:"column:device_id"`
	UserID             *int64  `gorm:"column:user_id"`
	UserName           *string `gorm:"column:username"`
	TotalPhones        int64   `gorm:"column:total_phones"`
	PendingCount       int64   `gorm:"column:pending_count"`
	CalledCount        int64   `gorm:"column:called_count"`
	InterestedCount    int64   `gorm:"column:interested_count"`
	NotInterestedCount int64   `gorm:"column:not_interested_count"`
}

// GetDeviceStats 获取所有设备的统计数据
func GetDeviceStats(ctx context.Context) ([]DeviceStatsRow, error) {
	var rows []DeviceStatsRow
	err := DB.WithContext(ctx).Raw(`
		SELECT
			ud.device_id,
			ud.user_id,
			u.username,
			COUNT(dp.id) AS total_phones,
			COALESCE(SUM(CASE WHEN dp.status = 0 THEN 1 ELSE 0 END), 0) AS pending_count,
			COALESCE(SUM(CASE WHEN dp.status = 1 THEN 1 ELSE 0 END), 0) AS called_count,
			COALESCE(SUM(CASE WHEN dp.status = 2 THEN 1 ELSE 0 END), 0) AS interested_count,
			COALESCE(SUM(CASE WHEN dp.status = 3 THEN 1 ELSE 0 END), 0) AS not_interested_count
		FROM user_device ud
		LEFT JOIN device_phone dp ON ud.device_id = dp.device_id AND dp.deleted_at IS NULL
		LEFT JOIN user u ON ud.user_id = u.id
		GROUP BY ud.device_id, ud.user_id, u.username
		ORDER BY ud.id ASC
	`).Scan(&rows).Error
	return rows, err
}
