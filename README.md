# CloudMind

项目结构：
app：所有的服务

common：通用组件

data： 存储数据

deploy：配置文件

golang：服务搭建文件夹

需要预先安装好 docker-compose 和 docker

项目启动顺序：
```text
1. docker network create CloudMind_net
建立网络，名为CloudMind_net
2. 运行容器启动脚本
bash run.sh
3. 创建数据库
通过Navicat连接mysql，创建数据库，然后创建对应的表，并导入数据
注意第一次使用需要配置mysql

docker exec -it mysql mysql -uroot -p
##输入密码：PXDN93VRKUm8TeE7
use mysql;
update user set host='%' where user='root';
FLUSH PRIVILEGES;


