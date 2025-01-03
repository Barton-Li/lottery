/*
 Navicat Premium Data Transfer

 Source Server         : lottery
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : usercenter

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 30/12/2024 17:13:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
    `id`                bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键，用户唯一标识符，自动递增',
    `create_time`       datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用户信息创建时间，默认为当前时间',
    `update_time`       datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '用户信息最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间',
    `delete_time`       datetime(0) DEFAULT NULL COMMENT '用户信息删除时间，默认为空',
    `del_state`         tinyint(0) NOT NULL DEFAULT 0 COMMENT '删除状态，0表示未删除，1表示已删除',
    `version`           bigint(0) NOT NULL DEFAULT 0 COMMENT '版本号，用于数据版本控制',
    `mobile`            char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用户手机号码，默认为空字符串',
    `password`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码，默认为空字符串',
    `nickname`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称，默认为空字符串',
    `sex`               tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别，0表示男，1表示女',
    `avatar`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '头像' COMMENT '用户头像，默认为"头像"字符串',
    `info`              varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户其他信息，默认为空字符串',
    `is_admin`          tinyint(1) DEFAULT 0 COMMENT '是否是管理员，1表示是管理员，0表示不是管理员',
    `signature`         varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '个性签名，默认为空字符串',
    `location_name`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '地址名称，默认为空字符串',
    `longitude`         DOUBLE PRECISION NOT NULL DEFAULT 0 COMMENT '经度，默认为0',
    `latitude`          DOUBLE PRECISION NOT NULL DEFAULT 0 COMMENT '纬度，默认为0',
    `total_prize`       int(0) NOT NULL DEFAULT 0 COMMENT '累计奖品数量，默认为0',
    `fans`              int(0) NOT NULL DEFAULT 0 COMMENT '粉丝数量，默认为0',
    `all_lottery`       int(0) NOT NULL DEFAULT 0 COMMENT '参与或发起的全部抽奖活动数量，默认为0',
    `initiation_record` int(0) NOT NULL DEFAULT 0 COMMENT '发起的抽奖活动数量，默认为0',
    `winning_record`    int(0) NOT NULL DEFAULT 0 COMMENT '中奖的记录数量，默认为0',
    PRIMARY KEY (`id`) USING BTREE COMMENT '主键索引，使用BTREE',
    UNIQUE INDEX `idx_mobile` (`mobile`) USING BTREE COMMENT '唯一索引，确保手机号码唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

-- ---------------------------
-- Table structure for user_auth
-- ---------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
    `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键，授权记录唯一标识符，自动递增',
    `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '授权记录创建时间，默认为当前时间',
    `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '授权记录最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间',
    `delete_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '授权记录删除时间，默认为当前时间',
    `del_state` tinyint(0) NOT NULL DEFAULT 0 COMMENT '删除状态，0表示未删除，1表示已删除',
    `version` bigint(0) NOT NULL DEFAULT 0 COMMENT '版本号，用于数据版本控制',
    `user_id` bigint(0) NOT NULL DEFAULT 0 COMMENT '用户ID，关联到user_info表中的id',
    `auth_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台唯一ID，用于标识用户在该平台上的唯一性',
    `auth_type` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台类型，标识授权的平台（如：微信、QQ等）',
    PRIMARY KEY (`id`) USING BTREE COMMENT '主键索引，使用BTREE',
    UNIQUE INDEX `idx_type_key` (`auth_type`, `auth_key`) USING BTREE COMMENT '唯一索引，确保每个平台类型的auth_key唯一',
    UNIQUE INDEX `idx_userID_key` (`user_id`, `auth_type`) USING BTREE COMMENT '唯一索引，确保每个用户在每个平台类型上只有一个授权记录'
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户授权表' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for user_address
-- ----------------------------
DROP TABLE IF EXISTS `user_address`;
CREATE TABLE `user_address` (
    `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键，地址记录唯一标识符，自动递增',
    `user_id` int NOT NULL DEFAULT '0' COMMENT '用户ID，关联到user_info表中的id',
    `contact_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人姓名，默认为空字符串',
    `contact_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系人手机号码，默认为空字符串',
    `district` json NOT NULL COMMENT '地区信息，以JSON格式存储',
    `detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址，默认为空字符串',
    `postcode` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮政编码，默认为空字符串',
    `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为默认地址，1表示是默认地址，0表示不是默认地址',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '地址记录创建时间，默认为当前时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '地址记录最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间',
    PRIMARY KEY (`id`) COMMENT '主键索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='用户收货地址表';


-- ---------------------------
-- Table structure for user_shop
-- ----------------------------
DROP  TABLE  IF EXISTS `user_shop`;
CREATE TABLE `user_shop` (
    `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键，店铺记录唯一标识符，自动递增',
    `user_id` int NOT NULL DEFAULT '0' COMMENT '用户ID，关联到user_info表中的id',
    `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称，默认为空字符串',
    `location` decimal(65,0) NOT NULL DEFAULT '0' COMMENT '店铺位置，以经纬度格式存储',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '店铺记录创建时间，默认为当前时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '店铺记录最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间',
    `delete_time` datetime DEFAULT NULL COMMENT '店铺记录删除时间，默认为空',
    PRIMARY KEY (`id`) USING  BTREE COMMENT '主键索引'
)ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户店铺表' ROW_FORMAT = Dynamic;
SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for user_sponsor
-- ---
-- ---
DROP TABLE IF EXISTS `user_sponsor`;
CREATE TABLE `user_sponsor` (
    `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键，推广记录唯一标识符，自动递增',
    `user_id` int(0) NOT NULL DEFAULT '0' COMMENT '用户ID，关联到user_info表中的id',
    `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1微信号 2公众号 3小程序 4微信群 5视频号',
    `applet_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT 'type=3时该字段才有意义，1小程序链接 2路径跳转 3二维码跳转',
   `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推广名称，默认为空字符串',
    `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推广描述，默认为空字符串',
    `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推广头像，默认为空字符串',
    `is_show` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否展示，1表示展示，2表示不展示',
    `qr_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '推广二维码，type=1 2 3&applet_type=3 4 默认为空字符串',
    `input_a` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=5 applet_type=2 or applet_type=1 推广输入框A，默认为空字符串',
    `input_b` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'type=5 applet_type=2  推广输入框B，默认为空字符串',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '推广记录创建时间，默认为当前时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '推广记录最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间',
    `delete_time` datetime DEFAULT NULL COMMENT '推广记录删除时间，默认为空',
    PRIMARY KEY (`id`) USING BTREE COMMENT '主键索引，使用BTREE'
)ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '抽奖发起人表' ROW_FORMAT = Dynamic;
SET  FOREIGN_KEY_CHECKS = 1;

-- ---------------------------
-- Table structure for user_contact
-- --------------------------------
DROP TABLE IF EXISTS `user_contact`;
CREATE TABLE `user_contact` (
    `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键，联系方式记录唯一标识符，自动递增',
    `user_id` int(0) NOT NULL DEFAULT '0' COMMENT '用户ID，关联到user_info表中的id',
    `content` json NOT NULL COMMENT '联系方式内容，以JSON格式存储',
    `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系方式备注，默认为空字符串',
        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '联系方式记录创建时间，默认为当前时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '联系方式记录最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间',
    PRIMARY KEY (`id`) USING BTREE COMMENT '主键索引，使用BTREE'

)ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户联系方式表' ROW_FORMAT = Dynamic;
SET FOREIGN_KEY_CHECKS = 1;

-- ---------------------------
-- Table structure for user_dynamic
-- ---
-- ---
DROP TABLE IF EXISTS `user_dynamic`;
CREATE TABLE `user_dynamic` (
    `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键，动态记录唯一标识符，自动递增',
    `user_id` int(0) NOT NULL DEFAULT '0' COMMENT '用户ID，关联到user_info表中的id',
    `dynamic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '动态链接，默认为空字符串',
    `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '动态备注，默认为空字符串',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '动态记录创建时间，默认为当前时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '动态记录最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间',
    `delete_time` datetime DEFAULT NULL COMMENT '动态记录删除时间，默认为空',
    PRIMARY KEY (`id`) USING BTREE COMMENT '主键索引，使用BTREE'
)ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户动态表' ROW_FORMAT = Dynamic;
SET FOREIGN_KEY_CHECKS = 1;

