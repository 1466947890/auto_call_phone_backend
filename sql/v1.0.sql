
-- 用户设备表
CREATE TABLE `user_device` (
    id INT PRIMARY KEY  AUTO_INCREMENT comment '主键ID',
    device_id varchar(128) NOT NULL comment '设备ID',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间'
)

-- 电话号码数据表

CREATE TABLE `device_phone` (
    id INT PRIMARY KEY AUTO_INCREMENT comment `主键ID`,
    phone_number varchar(128) NOT NULL comment `电话号码`,
    device_id varchar(128) NOT NULL comment `设备ID`,
    `status` tinyint(4) NOT NULL DEFAULT 0 comment `号码状态`,
    remarks  varchar(256) NOT NULL DEFAULT '' comment `备注`
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP comment `创建时间`,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment `更新时间`,
    UNIQUE KEY (`phone_number`, `device_id`)
)