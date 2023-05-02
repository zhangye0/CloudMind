SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `account` bigint NOT NULL AUTO_INCREMENT,
                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `del_state` tinyint NOT NULL DEFAULT '0',
                        `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                        `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        `sex` char(1) NOT NULL DEFAULT '0' COMMENT '性别 男, 女',
                        `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        PRIMARY KEY (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
                             `account` bigint NOT NULL AUTO_INCREMENT,
                             `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `del_state` tinyint NOT NULL DEFAULT '0',
                             `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                             `auth_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台唯一id',
                             `auth_type` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台类型',
                             PRIMARY KEY (`account`),
                             UNIQUE KEY `idx_type_key` (`auth_type`,`auth_key`) USING BTREE,
                             UNIQUE KEY `idx_userId_key` (`account`,`auth_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户授权表';

SET FOREIGN_KEY_CHECKS = 1;