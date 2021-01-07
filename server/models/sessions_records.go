package models

import "time"

/**
CREATE TABLE `sessions_records` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `session_id` int unsigned NOT NULL COMMENT '会话主键ID',
  `sender_id` int unsigned NOT NULL DEFAULT '0' COMMENT '发送者ID',
  `sender_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '发送者类型',
  `sender_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '发送者名称',
  `message` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '消息内容',
  `message_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '消息类型',
  `message_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '消息的发送时间',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `index_session_uid` (`session_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='记录会话的聊天信息表';
*/

type SessionsRecord struct {
	Model
	SessionId   *uint64
	SenderId    *uint64
	SenderType  *string
	SenderName  *string
	Message     *string
	MessageType *string
	MessageTime *time.Time
}
