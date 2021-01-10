package models_merchant

import "api/models"

/**
CREATE TABLE `tenants_department` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `parent_id` int unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `department_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '部门排序',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='租户部门表';
*/

type MerchantsDepartment struct {
	models.Model
	MerchantId     *uint64
	ParentId       *uint64
	DepartmentName *string
	Sort           *uint
	Desc           *string
}
