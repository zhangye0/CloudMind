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
| usercenter-api   | 2001 | 3002            |
| mqueue-job       | 2002 | 3004            |
| usercenter-rpc   | 4001 | 3003            |
| mqueue-scheduler | 4002 | 3005            |
| filecenter-api   | 2003 | 3006            |
| filecenter-rpc   | 4003 | 3007            |
| es-api           | 2004 | 3008            |             
| es-rpc           | 4004 | 3009            |
  

#### tips: API的端口从20开始, RPC的端口从40开始, Prometheus端口从30开始


Mysql :  自行客户端工具(datagrip)查看
- username : root
- 密码 : PXDN93VRKUm8TeE7

Redis :  自行工具（AnotherRedisDesktopManager）查看
- 密码 : G62m50oigInC30sf


## 三.Linux 环境配置

### 1. 下载vim
```
$ sudo apt install vim
```
### 2. 安装 Golang
```
(1). 下载地址： https://golang.google.cn/doc/install

// go1.20.4为例
(2). 解压压缩包至/usr/local :
   $ tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz

(3). 添加/usr/local/go/bin到环境变量 :
   $ vim $HOME/.profile
   $ export PATH=$PATH:/usr/local/go/bin
   $ source $HOME/.profile

(4). 验证安装结果
   $ go version
   go version go1.20.4 linux/amd64
```

### 3. 安装Goland
```
下载地址：https://www.jetbrains.com/go/download
完成后进入 Goland -> File -> Settings -> Go -> GOPATH 配置全局GOPATH // 选择一个文件夹即可
```

## 接下来的操作都在 Goland 终端操作
### 4. Go Module设置
```
(1). 查看GO111MODULE开启情况
   $ go env GO111MODULE
   on

(2). 开启GO111MODULE，如果已开启（即执行go env GO111MODULE结果为on）请跳过。
   $ go env -w GO111MODULE="on"

(3). 设置GOPROXY
   $ go env -w GOPROXY=https://goproxy.cn

(4). 设置GOMODCACHE

查看GOMODCACHE
$ go env GOMODCACHE

如果目录不为空或者/dev/null，请跳过。
$ go env -w GOMODCACHE=$GOPATH/pkg/mod
```

### 5. 安装 goctl
```
$ GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
```
注意安装完后，重启终端

### 6. protoc & protoc-gen-go安装
```
$ goctl env check -i -f --verbose
```

### 7. 安装配置 git
```
$ sudo apt-get install git
$ git config --global user.name Love-YeLin // 换成你的用户名
$ git config --global user.email 1807209079@qq.com // 换成你的邮箱

// 配置ssh 密钥
$ ssh-keygen -t rsa -b 4096 -C 1807209079@qq.com // 换成你的邮箱
$ cd .ssh
$ vim id_rsa.pub // 这里换成你生成的 .pub文件
复制密钥，然后去 github 创建密钥
```

### 8. 安装 docker
```
$ sudo apt install docker.io

// 配置用户组支持自定义请求的格式，默认的请求格式为 {subject, object, action}。
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
docker ps
sudo service docker restart     
sudo systemctl daemon-reload
sudo systemctl restart docker

// 安装docker-compose
下载地址：https://github.com/docker/compose/releases
// 将文件拖到主目录下，然后 cd 到主目录下
sudo mv docker-compose-linux-x86_64 /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```


## 四.系统功能

- [ ] usercenter
    - [ ] 登录
         - [x] 邮箱登录
         - [ ] QQ登录
         - [ ] 微信登录
         - [ ] ...
    - [x] 注册
    - [ ] 发送邮件
    - [x] JWT鉴权
    - [x] 退出登录
    - [x] 获取用户信息
    - [x] 修改用户信息
    - [ ] 实名认证
- [ ] mqueue
    - [x] 每月定时发放流量
- [ ] elasticsearch
    - [ ] 排行榜功能
      - [ ] 总/月/日榜功能
      - [x] 下载量/收藏量/点赞量的文件/帖子排行
    - [ ] 搜索功能
      - [ ] 搜索文件
        - [x] 按相关性排序
        - [ ] 按下载量排序
        - [ ] 按点赞量排序
        - [ ] 按收藏量排序
      - [ ] 搜索帖子
        - [x] 按相关性排序
        - [ ] 按点赞量排序
        - [ ] 按收藏量排序
        
    