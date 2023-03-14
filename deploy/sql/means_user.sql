/*
 Navicat MySQL Data Transfer

 Source Server         : means
 Source Server Type    : MySQL
 Source Server Version : 5.7
 Source Host           : 127.0.0.1:3306
 Source Schema         : user
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
                        `nickname` varchar(100) NOT NULL DEFAULT '' COMMENT '昵称',
                        `phone` varchar(20) NOT NULL COMMENT '手机号',
                        `sex` tinyint(1) DEFAULT '2' COMMENT '性别 1 男 2 女',
                        `avatar` text COMMENT '头像',
                        `password` varchar(255) NOT NULL COMMENT '密码',
                        `deleted_at` tinyint(1) DEFAULT NULL COMMENT '1 已删除',
                        `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                        `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;