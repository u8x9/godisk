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

