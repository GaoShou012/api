package models_statistics

import (
	"api/models"
	"time"
)

type StatsTenantsState struct {
	models.Model
	TenantsId   *uint64
	TenantsName *string
	MerchantId  *uint64
	CsGroup     *uint64
	State       *int
	StartAt     time.Time
	EndAt       time.Time
}

/*
CREATE TABLE `stats_tenants_state` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int unsigned DEFAULT NULL COMMENT '商户ID',
  `tenant_id` int unsigned DEFAULT NULL,
  `tenant_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '客服名',
  `cs_group` int DEFAULT NULL COMMENT '客服组ID',
  `state` tinyint DEFAULT NULL COMMENT '状态',
  `start_at` datetime DEFAULT NULL COMMENT '开始时间',
  `end_at` datetime DEFAULT NULL COMMENT '结束时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='客服状态统计';
*/
func (m *StatsTenantsState) GetTableName() string {
	return "stats_tenants_state"
}
