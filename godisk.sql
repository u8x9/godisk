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
