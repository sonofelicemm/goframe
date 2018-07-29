# sono
## 1 简介
sono项目是一个轻量级的go-web框架。开发者可以直接clone该仓库，将项目名sono替换为自己的模块名称，并将不同go文件中的sono替换掉即可。

## 2 主要功能列表
```
1. 静态路由
2. toml格式配置文件读取
3. 单元测试goCheck封装
4. log文件配置封装
5. mysql读取公共类封装
6. redis公共类封装
7. cron定时任务封装

```

## 3 使用说明
### 3.1 sono配置
toml文件进行mysql的链接信息、日志路径等相关的配置:

```
[log]
dir = "./logs"

[Mysql]
UserName = "sonofelice"
Password = "123456"
IpHost = "127.0.0.1:8902"
DbName = "sono_mysql"
```
### 3.1 sono工程启动

```
./sono ../conf/conf.toml
```
### 3.3 单测
框架集成了goCheck测试框架，完善了go单测的功能。