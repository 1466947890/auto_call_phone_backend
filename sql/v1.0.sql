
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

-- 用户表
CREATE TABLE `user` (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    username VARCHAR(64) NOT NULL COMMENT '用户名',
    password_hash VARCHAR(256) NOT NULL COMMENT '密码哈希',
    role VARCHAR(16) NOT NULL DEFAULT 'user' COMMENT '角色：user普通用户，admin管理员',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 给 user_device 表添加 user_id 字段和唯一索引
ALTER TABLE `user_device`
    ADD COLUMN user_id INT DEFAULT NULL COMMENT '用户ID' AFTER device_id,
    ADD INDEX `idx_user_device_user_id` (`user_id`),
    ADD UNIQUE KEY `uk_user_device` (`device_id`);

-- 给 user 表添加 role 字段
ALTER TABLE `user`
    ADD COLUMN role VARCHAR(16) NOT NULL DEFAULT 'user' COMMENT '角色：user普通用户，admin管理员' AFTER password_hash;

-- 插入默认管理员账号（用户名：admin，密码：123456）
INSERT INTO `user` (`username`, `password_hash`, `role`) VALUES ('admin', '$2a$10$xYEE6yK/UFU5Ke3xwnJmOO/m5hM1W0XxzGoxlSptvapi6T2h4mGa2', 'admin')
ON DUPLICATE KEY UPDATE `role` = 'admin';
