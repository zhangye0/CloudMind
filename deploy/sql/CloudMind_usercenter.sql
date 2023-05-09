SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
create table user
(
    id          bigint auto_increment
        primary key,
    email       varchar(255)                           not null comment '邮箱号',
    nickname    varchar(20)                            not null,
    password    varchar(255)                           not null,
    sex         tinyint(1)   default 0                 null comment '性别 0:男 1:女',
    avatar      varchar(255) default ''                null,
    name        varchar(30)  default ''                null,
    idcard      varchar(20)  default ''                null,
    create_time datetime     default CURRENT_TIMESTAMP null,
    update_time datetime     default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_time datetime     default CURRENT_TIMESTAMP null,
    del_state   tinyint      default 0                 null,
    constraint idx_mobile
        unique (email)
)
    comment '用户表';


-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
create table user_auth
(
    id          bigint auto_increment
        primary key,
    user_id     bigint      default 0                 not null,
    auth_key    varchar(64) default ''                not null comment '平台唯一id',
    auth_type   varchar(12) default ''                not null comment '平台类型',
    create_time datetime    default CURRENT_TIMESTAMP not null,
    update_time datetime    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    delete_time datetime    default CURRENT_TIMESTAMP not null,
    del_state   tinyint     default 0                 not null,
    constraint idx_type_key
        unique (auth_type, auth_key),
    constraint idx_userId_key
        unique (user_id, auth_type)
)
    comment '用户授权表';

