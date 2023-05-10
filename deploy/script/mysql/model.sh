#!/usr/bin/env bash

# 使用方法：
# bash model.sh usercenter user
# bash model.sh usercenter user_auth
# bash model.sh usercenter user_avatar

#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./model
# 数据库配置
host=127.0.0.1
port=33069
dbname=cloudmind_$1
username=root
passwd=PXDN93VRKUm8TeE7

echo "开始创建库：$dbname 的表：$2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero