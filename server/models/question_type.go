package models

/**
CREATE TABLE `question_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `category_id` int unsigned NOT NULL DEFAULT '0' COMMENT '所属大类ID',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '问题名称',
  `binding_setting` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '绑定设置 1所有客服 2 对话组 3 客服',
  `dialogue_group` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '对话组',
  `tenant_id` int unsigned NOT NULL COMMENT '客服',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `category_id` (`category_id`) USING BTREE,
  KEY `opt_id` (`merchant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='问题类型表';
*/

type QuestionType struct {
	Model
	CategoryId     *uint64
	MerchantId     *uint64
	Name           *string
	bindingSetting *int
	DialogueGroup  *string
	TenantId       *uint64
}
