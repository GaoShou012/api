/*
 Navicat Premium Data Transfer

 Source Server         : bob_kf
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : 127.0.0.1:13306
 Source Schema         : api

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 23/12/2020 18:46:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for rbac_role_assoc_api
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_assoc_api`;
CREATE TABLE `rbac_role_assoc_api` (
  `id` int NOT NULL,
  `tenant_id` int unsigned NOT NULL COMMENT '租户ID',
  `role_id` int unsigned NOT NULL COMMENT '角色ID',
  `api_id` int unsigned NOT NULL COMMENT 'API ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联API';

SET FOREIGN_KEY_CHECKS = 1;
