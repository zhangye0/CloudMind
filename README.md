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


# Linux 环境配置

### 一. 下载vim
```
$ sudo apt install vim
```
### 二. 安装 Golang
```
1. 下载地址： https://golang.google.cn/doc/install

// go1.15.8为例
2. 解压压缩包至/usr/local :
   $ tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz

3. 添加/usr/local/go/bin到环境变量 :
   $ vim $HOME/.profile
   $ /export PATH=$PATH:/usr/local/go/bin
   $ source $HOME/.profile

4. 验证安装结果
   $ go version
   go version go1.15.1 linux/amd64
```

### 三. 安装Goland
```
下载地址：https://www.jetbrains.com/go/download
完成后进入 Goland -> File -> Settings -> Go -> GOPATH 配置全局GOPATH // 选择一个文件夹即可
```

## 接下来的操作都在 Goland 终端操作
### 四. Go Module设置
```
1. 查看GO111MODULE开启情况
   $ go env GO111MODULE
   on

2. 开启GO111MODULE，如果已开启（即执行go env GO111MODULE结果为on）请跳过。
   $ go env -w GO111MODULE="on"

3. 设置GOPROXY
   $ go env -w GOPROXY=https://goproxy.cn

4. 设置GOMODCACHE

查看GOMODCACHE
$ go env GOMODCACHE

如果目录不为空或者/dev/null，请跳过。
$ go env -w GOMODCACHE=$GOPATH/pkg/mod
```

### 五. 安装 goctl
```
$ GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
```
注意安装完后，重启终端

### 六. protoc & protoc-gen-go安装
```
$ goctl env check -i -f --verbose
```

### 七. 安装配置 git
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

### 八. 安装 docker
```
$ sudo apt install docker.io

// 配置用户组
sudo groupadd docker
sudo usermod -aG docker $USER
sudo systemctl restart docker
docker ps
```
