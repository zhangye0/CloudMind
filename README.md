# CloudMind

项目结构：
- app：所有的服务

- common：通用组件

- data： 存储数据

- deploy：依赖的配置文件


#### tips: 需要预先安装好 docker-compose 和 docker

# 项目启动顺序：
1. docker network create CloudMind_net
建立网络，名为CloudMind_net

2.运行容器启动脚本
bash run.sh

3.创建数据库
通过Navicat连接mysql，创建数据库，然后创建对应的表，并导入数据

#### tips:注意第一次使用需要配置mysql
```shell
docker exec -it mysql mysql -uroot -p
##输入密码：PXDN93VRKUm8TeE7
use mysql;
update user set host='%' where user='root';
FLUSH PRIVILEGES;
```

| 依赖名   | 端口号  |
|-------|------|
| mysql | 33069 |
| redis | 6379 |
| kafka | 9092 |
| kibana | 9092 |
| Grafana | 3001 |
|jaeger | 16686 |
|asynq| 8980 |
| Prometheus | 9090 |
| Elastic search | 9200 |


Mysql :  自行客户端工具(datagrip)查看
- username : root
- 密码 : PXDN93VRKUm8TeE7

Redis :  自行工具（AnotherRedisDesktopManager）查看
- 密码 : G62m50oigInC30sf