package models_statistics

import "api/models"

type StatsOperateLog struct {
	models.Model
	OperateTenantId       *uint64 //操作者id
	OperateTenantName     *string //操作者名称
	OperateTenantRoleId   *uint64 //角色id
	OperateTenantRoleName *string //角色
	BeOperateTenantId     *uint64 //被操作者客服角色
	BeOperateTenantName   *string //被操作客服角色名
	OperateLog            *string //日志
	MerchantId            *uint64 //商户ID
	Params                *string //请求参数
}

/*
CREATE TABLE `stats_operate_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `operator_tenant_id` int unsigned DEFAULT NULL COMMENT '操作者id',
  `operator_tenant_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作者名称',
  `operator_tenant_role_id` int unsigned DEFAULT NULL COMMENT '角色id',
  `operator_tenant_role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色',
  `be_operator_tenant_id` int DEFAULT NULL COMMENT '被操作者客服角色',
  `be_operator_tenant_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '被操作客服角色',
  `operate_log` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '日志',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `params` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求参数',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='客服操作统计';
*/
func (m *StatsOperateLog) GetTableName() string {
	return "stats_operate_long"
}
