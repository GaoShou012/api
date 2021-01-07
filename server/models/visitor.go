package models

import "time"

/*
CREATE TABLE `visitor` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '唯一主键',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `customer_id` int unsigned NOT NULL COMMENT '访客ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客姓名',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客帐号',
  `level` int unsigned NOT NULL COMMENT '访客等级',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客标签',
  `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1为男性，0为女性，2未知',
  `phone` int unsigned NOT NULL DEFAULT '0' COMMENT '访客电话',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客邮箱',
  `wechat` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '微信号码',
  `wechat_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信昵称',
  `qq` int unsigned NOT NULL DEFAULT '0' COMMENT 'QQ号码',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '访客者IP地域',
  `creator` int unsigned DEFAULT '0' COMMENT '创建者ID',
  `modify_id` int unsigned DEFAULT '0' COMMENT '修改人ID',
  `sessions_times` int unsigned NOT NULL DEFAULT '0' COMMENT '会话次数',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `opt_id` (`merchant_id`,`customer_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='游客信息表';
*/
type Visitor struct {
	Model
	MerchantId   *uint64
	CustomerId   *uint64
	Name         *string
	Account      *string
	Level        *int
	Label        *string
	Gender       *int
	Phone        *uint64
	Email        *string
	Wechat       *string
	WechatName   *string
	QQ           *uint
	Ip           *string
	Creator      *uint
	ModifyId     *uint
	SessionsTime *time.Time
}
