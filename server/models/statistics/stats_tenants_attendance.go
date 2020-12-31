package models_statistics

import (
	"api/models"
	"time"
)

type StatTenantsAttendance struct {
	models.Model
	TenantsId   *uint64
	TenantsName *string
	MerchantId  *uint64
	ChatGroup   *uint64
	LoginTime   time.Time
	LogoutTime  time.Time
}

/*
CREATE TABLE `stats_tenants_attendance` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` int unsigned DEFAULT NULL COMMENT '客服ID',
  `tenant_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '客服姓名',
  `merchant_id` int unsigned DEFAULT NULL COMMENT '商户ID',
  `chat_group` int unsigned NOT NULL DEFAULT '0' COMMENT '对话组id',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `logout_time` datetime DEFAULT NULL COMMENT '退出时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='客服考勤统计';
*/
func (m *StatTenantsAttendance) GetTableName() string {
	return "stat_tenants_attendance"
}
