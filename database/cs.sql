/*
 Navicat Premium Data Transfer

 Source Server         : work_studio
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : 192.168.0.2:13306
 Source Schema         : cs

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 18/01/2021 17:02:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `enable` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否启用',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
  `roles` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admins
-- ----------------------------
BEGIN;
INSERT INTO `admins` VALUES (1, 1, 'admin', '$2a$10$rUrCtBkUe/QrUruCMbBAHe3AfVvobjgHVFUIFEobNH2angewlLuzy', '管理员', '1', '2021-01-14 22:32:49', '2021-01-14 22:32:49');
COMMIT;

-- ----------------------------
-- Table structure for admins_login_stats
-- ----------------------------
DROP TABLE IF EXISTS `admins_login_stats`;
CREATE TABLE `admins_login_stats` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `login_times` int unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `testing` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admins_login_stats
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ip_whitelist
-- ----------------------------
DROP TABLE IF EXISTS `ip_whitelist`;
CREATE TABLE `ip_whitelist` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP 地址',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `ip_group` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP 组',
  `last_login_time` datetime NOT NULL COMMENT 'ip最后登陆时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `opt_id` (`merchant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='黑名单表';

-- ----------------------------
-- Records of ip_whitelist
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchant
-- ----------------------------
DROP TABLE IF EXISTS `merchant`;
CREATE TABLE `merchant` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户号',
  `start_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '接入时间',
  `end_at` datetime NOT NULL COMMENT '到期时间',
  `channel` tinyint unsigned NOT NULL COMMENT '接入渠道 1 PC 2为 移动',
  `enable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 启用 1 禁用',
  `max_visitor` int unsigned NOT NULL DEFAULT '0' COMMENT '最大访客数',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='商户表';

-- ----------------------------
-- Records of merchant
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants
-- ----------------------------
DROP TABLE IF EXISTS `merchants`;
CREATE TABLE `merchants` (
  `id` int NOT NULL,
  `enable` tinyint(1) NOT NULL,
  `expiration` datetime NOT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of merchants
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_admins
-- ----------------------------
DROP TABLE IF EXISTS `merchants_admins`;
CREATE TABLE `merchants_admins` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `enable` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否启用',
  `state` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `user_type` tinyint unsigned NOT NULL COMMENT '用户类型',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of merchants_admins
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_admins_login_stats
-- ----------------------------
DROP TABLE IF EXISTS `merchants_admins_login_stats`;
CREATE TABLE `merchants_admins_login_stats` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `login_times` int unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `testing` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of merchants_admins_login_stats
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_department
-- ----------------------------
DROP TABLE IF EXISTS `merchants_department`;
CREATE TABLE `merchants_department` (
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

-- ----------------------------
-- Records of merchants_department
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_rbac_api
-- ----------------------------
DROP TABLE IF EXISTS `merchants_rbac_api`;
CREATE TABLE `merchants_rbac_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int unsigned NOT NULL COMMENT '租户ID',
  `method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'API 请求方式',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'API 请求路径',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统API';

-- ----------------------------
-- Records of merchants_rbac_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_rbac_menu
-- ----------------------------
DROP TABLE IF EXISTS `merchants_rbac_menu`;
CREATE TABLE `merchants_rbac_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int unsigned NOT NULL COMMENT '租户ID',
  `group_id` int unsigned NOT NULL COMMENT '菜单组ID',
  `sort` int unsigned NOT NULL COMMENT '排序',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单编码，关联前端',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单描述',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单项（二级菜单）';

-- ----------------------------
-- Records of merchants_rbac_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_rbac_menu_group
-- ----------------------------
DROP TABLE IF EXISTS `merchants_rbac_menu_group`;
CREATE TABLE `merchants_rbac_menu_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int unsigned NOT NULL COMMENT '租户ID',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单编码',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单组名字',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单组图标',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单组描述',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单组（一级菜单）';

-- ----------------------------
-- Records of merchants_rbac_menu_group
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_rbac_role
-- ----------------------------
DROP TABLE IF EXISTS `merchants_rbac_role`;
CREATE TABLE `merchants_rbac_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int unsigned NOT NULL COMMENT '租户ID',
  `enable` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否启用，0=禁用，1=启用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名字',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色图标',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色描述',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- ----------------------------
-- Records of merchants_rbac_role
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_rbac_role_assoc_api
-- ----------------------------
DROP TABLE IF EXISTS `merchants_rbac_role_assoc_api`;
CREATE TABLE `merchants_rbac_role_assoc_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL COMMENT '角色ID',
  `api_id` int unsigned NOT NULL COMMENT 'API ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`role_id`,`api_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联API';

-- ----------------------------
-- Records of merchants_rbac_role_assoc_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_rbac_role_assoc_menu
-- ----------------------------
DROP TABLE IF EXISTS `merchants_rbac_role_assoc_menu`;
CREATE TABLE `merchants_rbac_role_assoc_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL COMMENT '角色ID',
  `menu_id` int unsigned NOT NULL COMMENT '菜单ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`role_id`,`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联菜单（二级菜单）';

-- ----------------------------
-- Records of merchants_rbac_role_assoc_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants_rbac_role_assoc_menu_group
-- ----------------------------
DROP TABLE IF EXISTS `merchants_rbac_role_assoc_menu_group`;
CREATE TABLE `merchants_rbac_role_assoc_menu_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL COMMENT '角色ID',
  `menu_group_id` int unsigned NOT NULL COMMENT '菜单组ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`role_id`,`menu_group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联菜单组（一级菜单）';

-- ----------------------------
-- Records of merchants_rbac_role_assoc_menu_group
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for question_type
-- ----------------------------
DROP TABLE IF EXISTS `question_type`;
CREATE TABLE `question_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `category_id` int unsigned NOT NULL DEFAULT '0' COMMENT '所属大类ID',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '问题名称',
  `binding_setting` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '绑定设置 1所有客服 2 对话组 3 客服',
  `dialogue_group` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '对话组',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `category_id` (`category_id`) USING BTREE,
  KEY `opt_id` (`merchant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='问题类型表';

-- ----------------------------
-- Records of question_type
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for rbac_api
-- ----------------------------
DROP TABLE IF EXISTS `rbac_api`;
CREATE TABLE `rbac_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'API 请求方式',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'API 请求路径',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`method`,`path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=337 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统API';

-- ----------------------------
-- Records of rbac_api
-- ----------------------------
BEGIN;
INSERT INTO `rbac_api` VALUES (1, 'POST', '/admin/v1/login', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (2, 'POST', '/admin/v1/register', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (3, 'GET', '/admin/v1/auth_code', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (4, 'GET', '/admin/v1/logout', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (5, 'GET', '/admin/v1/operateor/info', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (6, 'POST', '/admin/v1/rbac/api/create', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (7, 'POST', '/admin/v1/rbac/api/update', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (8, 'GET', '/admin/v1/rbac/api/delete', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (9, 'POST', '/admin/v1/rbac/menu/create', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (10, 'POST', '/admin/v1/rbac/menu/update', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (11, 'GET', '/admin/v1/rbac/menu/delete', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (12, 'POST', '/admin/v1/rbac/role/create', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (13, 'POST', '/admin/v1/rbac/role/update', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (14, 'GET', '/admin/v1/rbac/role/delete', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (15, 'GET', '/admin/v1/rbac/role/select', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (16, 'POST', '/admin/v1/rbac/menu_group/create', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (17, 'POST', '/admin/v1/rbac/menu_group/update', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (18, 'GET', '/admin/v1/rbac/menu_group/delete', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (19, 'POST', '/admin/v1/rbac/role_assoc_api/create', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_api` VALUES (20, 'GET', '/admin/v1/rbac/role_assoc_api/delete', '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_api` VALUES (21, 'POST', '/admin/v1/rbac/role_assoc_menu_group/create', '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_api` VALUES (22, 'GET', '/admin/v1/rbac/role_assoc_menu_group/delete', '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_api` VALUES (23, 'POST', '/admin/v1/rbac/role_assoc_menu/create', '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_api` VALUES (24, 'GET', '/admin/v1/rbac/role_assoc_menu/delete', '2021-01-14 21:43:31', '2021-01-14 21:43:31');
COMMIT;

-- ----------------------------
-- Table structure for rbac_menu
-- ----------------------------
DROP TABLE IF EXISTS `rbac_menu`;
CREATE TABLE `rbac_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int unsigned NOT NULL COMMENT '菜单组ID',
  `sort` int unsigned NOT NULL COMMENT '排序',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单编码，关联前端',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单描述',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`group_id`,`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单项（二级菜单）';

-- ----------------------------
-- Records of rbac_menu
-- ----------------------------
BEGIN;
INSERT INTO `rbac_menu` VALUES (1, 1, 1, 'role', '角色管理', 'icon', 'desc', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_menu` VALUES (2, 1, 2, 'menu', '菜单管理', 'icon', 'desc', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
INSERT INTO `rbac_menu` VALUES (3, 1, 3, 'api', '接口管理', 'icon', 'desc', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
COMMIT;

-- ----------------------------
-- Table structure for rbac_menu_group
-- ----------------------------
DROP TABLE IF EXISTS `rbac_menu_group`;
CREATE TABLE `rbac_menu_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单组名字',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单组图标',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单组描述',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单组（一级菜单）';

-- ----------------------------
-- Records of rbac_menu_group
-- ----------------------------
BEGIN;
INSERT INTO `rbac_menu_group` VALUES (1, 1, '菜单与权限', '', '', '2021-01-14 21:43:30', '2021-01-14 21:43:30');
COMMIT;

-- ----------------------------
-- Table structure for rbac_role
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role`;
CREATE TABLE `rbac_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `enable` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否启用，0=禁用，1=启用',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名字',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色图标',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色描述',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- ----------------------------
-- Records of rbac_role
-- ----------------------------
BEGIN;
INSERT INTO `rbac_role` VALUES (1, 1, '超级管理员', '', '超级管理员', '2021-01-14 14:20:47', '2021-01-14 22:10:24');
COMMIT;

-- ----------------------------
-- Table structure for rbac_role_assoc_api
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_assoc_api`;
CREATE TABLE `rbac_role_assoc_api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL COMMENT '角色ID',
  `api_id` int unsigned NOT NULL COMMENT 'API ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`role_id`,`api_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=313 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联API';

-- ----------------------------
-- Records of rbac_role_assoc_api
-- ----------------------------
BEGIN;
INSERT INTO `rbac_role_assoc_api` VALUES (1, 1, 1, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (2, 1, 2, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (3, 1, 3, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (4, 1, 4, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (5, 1, 5, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (6, 1, 6, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (7, 1, 7, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (8, 1, 8, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (9, 1, 9, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (10, 1, 10, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (11, 1, 11, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (12, 1, 12, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (13, 1, 13, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (14, 1, 14, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (15, 1, 15, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (16, 1, 16, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (17, 1, 17, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (18, 1, 18, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (19, 1, 19, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (20, 1, 20, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (21, 1, 21, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (22, 1, 22, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (23, 1, 23, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_api` VALUES (24, 1, 24, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
COMMIT;

-- ----------------------------
-- Table structure for rbac_role_assoc_menu
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role_assoc_menu`;
CREATE TABLE `rbac_role_assoc_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL COMMENT '角色ID',
  `menu_id` int unsigned NOT NULL COMMENT '菜单ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UIDX` (`role_id`,`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联菜单（二级菜单）';

-- ----------------------------
-- Records of rbac_role_assoc_menu
-- ----------------------------
BEGIN;
INSERT INTO `rbac_role_assoc_menu` VALUES (1, 1, 1, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_menu` VALUES (2, 1, 2, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
INSERT INTO `rbac_role_assoc_menu` VALUES (3, 1, 3, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色关联菜单组（一级菜单）';

-- ----------------------------
-- Records of rbac_role_assoc_menu_group
-- ----------------------------
BEGIN;
INSERT INTO `rbac_role_assoc_menu_group` VALUES (1, 1, 1, '2021-01-14 21:43:31', '2021-01-14 21:43:31');
COMMIT;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `state` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0派对中，1会话处理中，2会话结束',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `session_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '会话ID',
  `user_id` int unsigned NOT NULL COMMENT '访客ID',
  `user_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客来源url',
  `user_vip_level` int unsigned NOT NULL COMMENT '访客等级',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '访客名称',
  `user_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '访客IP',
  `user_device` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1.PC;2.手机',
  `user_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客会话token',
  `user_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客地理位置',
  `user_rating` int unsigned NOT NULL DEFAULT '0' COMMENT '访客评分 0未评分，1非常满意，2比较满意，3一般，4不满意',
  `user_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客意见',
  `user_rating_time` datetime DEFAULT NULL COMMENT '访客评价时间',
  `cs_id` int unsigned NOT NULL DEFAULT '0' COMMENT '客服ID',
  `cs_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客服名称',
  `cs_group` int unsigned NOT NULL DEFAULT '0' COMMENT '客服组标识id',
  `cs_department` int unsigned NOT NULL DEFAULT '0' COMMENT '客服行政部门id',
  `service_tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务标签，归档的时候，',
  `service_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务类型',
  `service_topic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务主题',
  `cs_value` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '客服评估:1待定评价，2无价值，3有价值，4很有价值，5价值待定',
  `cs_comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客服评价备注',
  `service_request_time` datetime DEFAULT NULL COMMENT '访客请求服务的时间,等于创建服务时间',
  `service_begin_time` datetime DEFAULT NULL COMMENT '服务开始时间',
  `service_end_time` datetime DEFAULT NULL COMMENT '服务结束时间',
  `service_end_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '服务结束原因',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `session_id` (`session_id`) USING BTREE COMMENT '会话uuid',
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='会话列表';

-- ----------------------------
-- Records of sessions
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sessions_records
-- ----------------------------
DROP TABLE IF EXISTS `sessions_records`;
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

-- ----------------------------
-- Records of sessions_records
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for stats_merchants_attendance
-- ----------------------------
DROP TABLE IF EXISTS `stats_merchants_attendance`;
CREATE TABLE `stats_merchants_attendance` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `merchant_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '客服姓名',
  `merchant_id` int unsigned DEFAULT NULL COMMENT '商户ID',
  `chat_group` int unsigned NOT NULL DEFAULT '0' COMMENT '对话组id',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `logout_time` datetime DEFAULT NULL COMMENT '退出时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='客服考勤统计';

-- ----------------------------
-- Records of stats_merchants_attendance
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for stats_merchants_state
-- ----------------------------
DROP TABLE IF EXISTS `stats_merchants_state`;
CREATE TABLE `stats_merchants_state` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int unsigned DEFAULT NULL COMMENT '商户ID',
  `merchant_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '客服名',
  `cs_group` int DEFAULT NULL COMMENT '客服组ID',
  `state` tinyint DEFAULT NULL COMMENT '状态',
  `start_at` datetime DEFAULT NULL COMMENT '开始时间',
  `end_at` datetime DEFAULT NULL COMMENT '结束时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='客服状态统计';

-- ----------------------------
-- Records of stats_merchants_state
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for stats_operate_log
-- ----------------------------
DROP TABLE IF EXISTS `stats_operate_log`;
CREATE TABLE `stats_operate_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `operator_merchant_id` int unsigned DEFAULT NULL COMMENT '操作者id',
  `operator_merchant_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作者名称',
  `operator_merchant_role_id` int unsigned DEFAULT NULL COMMENT '角色id',
  `operator_merchant_role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色',
  `be_operator_merchant_id` int DEFAULT NULL COMMENT '被操作者客服角色',
  `be_operator_merchant_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '被操作客服角色',
  `operate_log` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '日志',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `params` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求参数',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='客服操作统计';

-- ----------------------------
-- Records of stats_operate_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for visitor
-- ----------------------------
DROP TABLE IF EXISTS `visitor`;
CREATE TABLE `visitor` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '唯一主键',
  `merchant_id` int unsigned NOT NULL DEFAULT '0' COMMENT '商户ID',
  `customer_id` int unsigned NOT NULL COMMENT '访客ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客姓名',
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客帐号',
  `level` int unsigned NOT NULL COMMENT '访客等级',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客标签',
  `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1为男性，0为女性，2未知',
  `phone` int unsigned NOT NULL DEFAULT '0' COMMENT '访客电话',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访客邮箱',
  `wechat` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '微信号码',
  `wechat_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '微信昵称',
  `qq` int unsigned NOT NULL DEFAULT '0' COMMENT 'QQ号码',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '访客者IP地域',
  `creator` int unsigned DEFAULT '0' COMMENT '创建者ID',
  `modify_id` int unsigned DEFAULT '0' COMMENT '修改人ID',
  `sessions_times` int unsigned NOT NULL DEFAULT '0' COMMENT '会话次数',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `opt_id` (`merchant_id`,`customer_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='游客信息表';

-- ----------------------------
-- Records of visitor
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
