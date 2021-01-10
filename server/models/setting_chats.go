package models

//对话设置
type SettingChats struct {
	Model
	AssignRule          *int    //分配规则 1按饱和度分配，2依次分配
	IsAssignLastService *bool   //是否优先分配上一次对话过的客服
	AutoRespond         *string //自动回复
	Welcome             *string //欢迎语
}

func (m *SettingChats) GetTableName() string {
	return "setting_chats"
}

/*
CREATE TABLE `setting_dialogue` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `opt_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `auto_rule` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '分配规则 1 包和度分配 2 依次分配',
  `is_priority` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '优先分配给上一次对话的客服 0是 1 否',
  `auto_answer` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '公司统一设置 0是 1 否',
  `auto_welcome` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '自动应答欢迎语',
  `answer_time` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '自动应答时间(秒)',
  `reply_times` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '同一访客最多自动回复次数',
  `auto_reply` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '自动回复话术',
  `auto_reply_two` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '自动回复话术2',
  `is_queue_tips` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '高排队请求人工限制提醒 0 不提示 1 提示',
  `queue_number` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '当排队人数超过多少位提示',
  `tips_visitor` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '提示访客',
  `is_send_tips` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '排队发消息提醒 0 不提示 1 提示',
  `queue_tips` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '当访客在排队时，发消息提示访客',
  `is_limit` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '请求人工限制 0 不限制 1 限制',
  `request_time` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '访客在几分钟内请求',
  `service_times` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '请求人工服务次数',
  `no_request_time` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '几分钟内不能再次请求',
  `tips_limit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '并提示访客',
  `is_no_response` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '访客无应答提示 0 不提示 1 提示',
  `no_response_time` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '访客在几秒内无应答',
  `response_tips` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发送消息提示访客',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `opt_id` (`opt_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='系统中的对话设置';
*/
