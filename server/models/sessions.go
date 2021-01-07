package models

import "time"

/**
CREATE TABLE `sessions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0派对中，1会话处理中，2会话结束',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `session_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '会话ID',
  `user_id` int unsigned NOT NULL COMMENT '访客ID',
  `user_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客来源url',
  `user_vip_level` int unsigned NOT NULL COMMENT '访客等级',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '访客名称',
  `user_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '访客IP',
  `user_device` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1.PC;2.手机',
  `user_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客会话token',
  `user_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客地理位置',
  `user_rating` int unsigned NOT NULL DEFAULT '0' COMMENT '访客评分 0未评分，1非常满意，2比较满意，3一般，4不满意',
  `user_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客意见',
  `user_rating_time` datetime DEFAULT NULL COMMENT '访客评价时间',
  `cs_id` int unsigned NOT NULL DEFAULT '0' COMMENT '客服ID',
  `cs_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客服名称',
  `cs_group` int unsigned NOT NULL DEFAULT '0' COMMENT '客服组标识id',
  `cs_department` int unsigned NOT NULL DEFAULT '0' COMMENT '客服行政部门id',
  `service_tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务标签，归档的时候，',
  `service_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务类型',
  `service_topic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务主题',
  `cs_value` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '客服评估:1待定评价，2无价值，3有价值，4很有价值，5价值待定',
  `cs_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客服评价备注',
  `service_request_time` datetime DEFAULT NULL COMMENT '访客请求服务的时间,等于创建服务时间',
  `service_begin_time` datetime DEFAULT NULL COMMENT '服务开始时间',
  `service_end_time` datetime DEFAULT NULL COMMENT '服务结束时间',
  `service_end_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务结束原因',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `session_id` (`session_id`) USING BTREE COMMENT '会话uuid',
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='会话列表';
*/
type Sessions struct {
	Status             *int
	MerchantId         *uint64
	SessionId          *string
	UserId             *uint64
	UserSource         *string
	UserVipLevel       *int
	UserName           *string
	UserIp             *string
	UserDevice         *int
	UserToken          *string
	UserLocation       *string
	UserRating         *int64
	UserComment        *string
	UserRatingTime     *time.Time
	CsId               *uint64
	CsName             *string
	CsGroup            *int
	CsDepartment       *int
	ServiceTags        *string
	ServiceType        *string
	ServiceTopic       *string
	CsValue            *int
	CsComment          *string
	ServiceRequestTime *time.Time
	ServiceBeginTie    *time.Time
	ServiceEndTime     *time.Time
	ServiceEndReason   *string
}
