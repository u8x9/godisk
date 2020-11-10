CREATE TABLE `tbl_file`(
    `id` int not null auto_increment,
    `file_sha1` char(48) not null default '' comment '文件hash',
    `file_name` varchar(256) not null default '' comment '文件名',
    `file_size` bigint default 0 comment '文件大小',
    `file_addr` varchar(1024) not null default '' comment '文件存储位置', 
    `create_at` datetime default now() comment '创建日期',
    `update_at` datetime default now() on update current_timestamp() comment '更新日期',
    `status` int not null default 0 comment '状态(可用/禁用/已删除)',
    `ext1` int default 0 comment '备用字段1',
    `ext2` text comment '备用字段2',
    primary key (`id`),
    unique key `idx_file_hash` (`file_sha1`),
    key `idx_status` (`status`)
) ENGINE=INNODB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE tbl_user(
    id int not null auto_increment,
    username varchar(64) not null comment '用户名',
    password varchar(256) not null comment '密码',
    email varchar(64) not null default '' comment '邮箱',
    phone char(11) not null default '' comment '手机号码',
    email_validated tinyint default 0 comment '邮箱是否验证',
    phone_validated tinyint default 0 comment '手机是否验证',
    signup_at datetime default current_timestamp comment '注册时间',
    last_active datetime default current_timestamp on update current_timestamp comment '最后活跃时间',
    profile text comment '用户属性',
    `status` tinyint not null default 0 comment '用户状态(启用/禁用/锁定/标记删除)',
    primary key (id),
    unique key `idx_phone` (phone),
    key `idx_status` (`status`)
) engine=INNODB default charset=UTF8MB4;

CREATE TABLE tbl_user_token(
    id int not null auto_increment,
    username varchar(64) not null,
    token char(40) not null,
    primary key(id),
    unique key `idx_username` (username)
) engine=INNODB default charset=utf8mb4;
