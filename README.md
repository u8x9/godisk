# Go实战仿百度云盘 实现企业级分布式云存储系统

## 服务架构

![image-20201103005708587](image-20201103005708587.png)

## 接口列表

| 接口             | url                  |
| ---------------- | -------------------- |
| 文件上传         | POST `/file/upload`  |
| 文件查询         | GET `/file/query`    |
| 文件下载         | GET `/file/download` |
| 文件删除         | POST `/file/delete`  |
| 文件修改(重命名) | POST `/file/update`  |

## 架构变迁(ch03)

![image-20201107163424732](image-20201107163424732.png)

## MYSQL主从

### 容器

```bash
docker network create godisk && \
docker run --name godisk_mysql_master --net godisk -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql --server-id=1 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci && \
docker run --name godisk_mysql_slave --net godisk -e MYSQL_ROOT_PASSWORD=root -p 3307:3306 -d mysql --server-id=2 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

### master

```bash
docker exec -it godisk_mysql_master mysql -uroot -p
```

```mysql
SHOW MASTER STATUS;
```

### slave

> 1. 下面的`MASTER_LOG_FILE`的取值来自于 master 节点上 `SHOW MASTER STATUS` 的结果
> 2. `MASTER_HOST` 是 master 节点容器的ip, 可通过 `docker inspect  mysql_master`查看
> 3. 实际应用时，避免使用`root`用户

```mysql
CHANGE MASTER TO MASTER_HOST='172.29.0.2', MASTER_USER='root', MASTER_PASSWORD='root', MASTER_LOG_FILE='binlog.000002', MASTER_LOG_POS=0, get_master_public_key=1;

START SLAVE;

SHOW SLAVE STATUS\G
```

## mysql 分库分表

### 水平分表

假设分成256张文件表，按文件hash后两位来切分，则以`tbl_${file_hash}[:-2]`的规则到对应表进行存取。

### 垂直分表

将表的字段拆分为不同的表。

