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

 Date: 23/12/2020 18:46:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for rbac_menu_group
-- ----------------------------
DROP TABLE IF EXISTS `rbac_menu_group`;
CREATE TABLE `rbac_menu_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` int unsigned NOT NULL COMMENT '租户ID',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单编码',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单组名字',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单组图标',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单组描述',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单组（一级菜单）';

SET FOREIGN_KEY_CHECKS = 1;
