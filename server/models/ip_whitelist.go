package models

import "time"

/**
CREATE TABLE `ip_whitelist` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP 地址',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `ip_group` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP 组',
  `last_login_time` datetime NOT NULL COMMENT 'ip最后登陆时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `opt_id` (`merchant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='黑名单表';
*/

type IpWhitelist struct {
	Model
	MerchantId    *uint
	Ip            *string
	Desc          *string
	IpGroup       *string
	LastLoginTime *time.Time
}
