# CloudMind

## 一.项目结构：
- app：所有的服务

- common：通用组件

- data： 存储数据

- deploy：依赖的配置文件

- docker-compose.yml： 项目启动的配置文件

- modd.conf： modd热加载插件的配置文件

#### tips: 需要预先安装好 docker-compose 和 docker

## 二.项目启动顺序：
1.建立网络，名为CloudMind_net

```shell
docker network create CloudMind_net
```

2.运行容器启动脚本
```shell
bash run.sh
```
##### tips: 如果遇到elasticsearch启动不了， 那就输入 
```shell
sudo chown -R $USER:$USER data/elasticsearch
```

3.创建数据库

通过 deploy/sql 目录下的sql语句直接导入数据库进行建表， 可通过DataGrip等可视化工具实现。

#### tips:注意第一次使用需要配置mysql
```shell
docker exec -it mysql mysql -uroot -p
##输入密码：PXDN93VRKUm8TeE7
use mysql;
update user set host='%' where user='root';
FLUSH PRIVILEGES;
```

| 依赖名            | 端口号   |
|----------------|-------|
| mysql          | 33069 |
| redis          | 6379  |
| kafka          | 9092  |
| kibana         | 9092  |
| Grafana        | 3001  |
| jaeger         | 16686 |
| asynq          | 8980  |
| Prometheus     | 9090  |
| Elastic search | 9200  |


| 类型名              | 端口号  | Prometheus监听端口号 | 
|------------------|------|-----------------|
| usercenter-api   | 2001 | 3001            |
| mqueue-job       | 2002 | 3003            |
| usercenter-rpc   | 4001 | 3002            |
| mqueue-scheduler | 4002 | 3004            |

#### tips: API的端口从20开始, RPC的端口从40开始, Prometheus端口从30开始


Mysql :  自行客户端工具(datagrip)查看
- username : root
- 密码 : PXDN93VRKUm8TeE7

Redis :  自行工具（AnotherRedisDesktopManager）查看
- 密码 : G62m50oigInC30sf