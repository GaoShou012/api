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

 Date: 26/12/2020 12:41:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for rbac_role_assoc_menu_group
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_assoc_menu_group`;
CREATE TABLE `rbac_role_assoc_menu_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL COMMENT '角色ID',
  `menu_group_id` int unsigned NOT NULL COMMENT '菜单组ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`role_id`,`menu_group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联菜单组（一级菜单）';

SET FOREIGN_KEY_CHECKS = 1;
