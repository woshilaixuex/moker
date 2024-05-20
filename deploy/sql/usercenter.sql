SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user_salt`;
CREATE TABLE `user_salt`(
                            `id` bigint NOT NULL AUTO_INCREMENT,
                            `user_id` bigint NOT NULL,
                            `salt` varchar(128) NOT NULL COMMENT '加密盐',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户加密表';

DROP TABLE IF EXISTS `users`;
CREATE TABLE `user` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `del_state` tinyint NOT NULL DEFAULT '0',
                        `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                        `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        `sex` tinyint(1) NOT NULL DEFAULT '3' COMMENT '性别  0:男  1:女  3:未知',
                        `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `role` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `user_id` bigint NOT NULL,
                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `del_state` tinyint NOT NULL DEFAULT '0',
                        `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                        `role_name` varchar(32) NOT NULL,
                        `role_zh_name` varchar(32) NOT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_role_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户权限表';

DROP TABLE IF EXISTS `students`;
CREATE TABLE `students` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `user_id` bigint NOT NULL,
                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `del_state` tinyint NOT NULL DEFAULT '0',
                        `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                        `name` varchar(32) NOT NULL DEFAULT '',
                        `major` varchar(32) NOT NULL DEFAULT '',
                        `faculty` varchar(32) NOT NULL DEFAULT '',
                        `school` varchar(32) NOT NULL DEFAULT '',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_suser_id` (`id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='学生表';

DROP TABLE IF EXISTS `teachers`;
CREATE TABLE `teachers` (
                            `id` bigint NOT NULL AUTO_INCREMENT,
                            `user_id` bigint NOT NULL,
                            `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `del_state` tinyint NOT NULL DEFAULT '0',
                            `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
                            `name` varchar(32) NOT NULL DEFAULT '',
                            `faculty` varchar(32) NOT NULL DEFAULT '',
                            `school` varchar(32) NOT NULL DEFAULT '',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `idx_tuser_id` (`id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='教师表';