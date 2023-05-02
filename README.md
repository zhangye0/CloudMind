# CloudMind

项目结构：
app：所有的服务

common：通用组件

data： 存储数据

deploy：配置文件

golang：服务搭建文件夹


项目启动顺序：
```text
1. docker network create CloudMind_net
建立网络，名为CloudMind_net
2. docker-compose -f docker-compose-env.yml up -d 
运行服务
3. docker-compose -f docker-compose-env.yml ps
查看依赖运行状态， up为正常， Restart为异常
4. 导入数据库
docker exec -it mysql mysql -uroot -p
输入密码：PXDN93VRKUm8TeE7
use mysql;
update user set host='%' where user='root';
FLUSH PRIVILEGES;
5. 创建数据库
通过Navicat连接mysql，创建数据库，然后创建对应的表，并导入数据
6. 启动项目
docker-compose up -d 
7. 进入项目运行服务
docker exec -it cloudmind_golang_1 bash
再运行对应的api和rpc即可
