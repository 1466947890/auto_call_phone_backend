
-- 用户设备表
CREATE TABLE `user_device` (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    device_id VARCHAR(128) NOT NULL COMMENT '设备ID',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
);

-- 电话号码数据表
CREATE TABLE `device_phone` (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    phone_number VARCHAR(128) NOT NULL COMMENT '电话号码',
    device_id VARCHAR(128) NOT NULL COMMENT '设备ID',
    `status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '号码状态',
    remarks VARCHAR(256) NOT NULL DEFAULT '' COMMENT '备注',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME NULL DEFAULT NULL COMMENT '软删除时间',
    UNIQUE KEY `uk_phone_device` (`phone_number`, `device_id`),
    KEY `idx_device_phone_deleted_at` (`deleted_at`)
);
