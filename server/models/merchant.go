package models

import (
	"api/global"
	"time"
)

/**
CREATE TABLE `merchant` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户号',
  `start_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '接入时间',
  `end_at` datetime NOT NULL COMMENT '到期时间',
  `channel` tinyint unsigned NOT NULL COMMENT '接入渠道 1 PC 2为 移动',
  `enable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 启用 1 禁用',
  `max_visitor` int unsigned NOT NULL DEFAULT '0' COMMENT '最大访客数',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='商户表';
*/

type Merchants struct {
	Model
	Name       *string
	Code       *string
	StartAt    *time.Time
	EndAt      *time.Time
	Channel    *int
	enable     *bool
	MaxVisitor *int
	Desc       *string
}

func (m *Merchants) GetTableName() string {
	return "merchants"
}

func (m *Merchants) SelectByCode(fields string, code string) (bool, error) {
	res := global.DBSlave.Table(m.GetTableName()).Select(fields).Where("code=?", code).First(m)
	if res.Error != nil {
		if res.RecordNotFound() {
			return false, nil
		} else {
			return false, res.Error
		}
	}
	return true, nil
}

// 商户是否启用
func (m *Merchants) IsEnable() bool {

}

// 商户是否租约过期
func (m *Merchants) IsExpiration() bool {

}
