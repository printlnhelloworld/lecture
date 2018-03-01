# 线上讲座项目

## 目录说明

* devdocs 文档目录
  * 原型
  * 接口
  * 流程图
* src 源码
  * backend 后端源码
  * frontend 前端源码
* test 集成测试

## Golang Tools

使用 [dep](https://github.com/golang/dep) 做包管理，后台的所有源码都在 `src/backend/` 里面，其中的 `vendor` 是后端依赖目录。不会包含到目录里面

* 采用 [Gin](https://github.com/gin-gonic/gin) 做框架
* 使用 [packr](https://github.com/gobuffalo/packr) 来包含前端生成文件，达到单文件发行的目的。
* 使用 [grom](github.com/jinzhu/gorm) 操作 MySQL 数据库
* 使用 [go.uuid](github.com/satori/go.uuid) 来生成用户 Token
* 使用 [go-toml](github.com/pelletier/go-toml) 来解析配置文件

## 版本控制

使用 Git 做版本管理

## 怎么编译运行

首先需要安装 dep、go 、packr、git，同时需要配置好 git、dep、go 等的代理设置，防止有些包没有办法获取到。

进入 `src/backend/`， 运行 `dep ensure` 进行依赖安装 ，`packr build` 会将前端文档编译包含的编译好的二进制中。然后复制 `conf/app.toml.example` 为 `conf/app.toml`，同时配置监听端口，监听路径、以及配置好的配置文件，然后运行即可。

对于使用 nginx 反代的情况，一个参考配置

```conf
server {
    listen 443 ssl http2;
    server_name lecture.hduhelp.com;

    access_log /var/log/nginx/lecture.hduhelp.com-access.log;

    ssl_certificate  /etc/letsencrypt/live/lecture.hduhelp.com/fullchain.pem;
    ssl_certificate_key  /etc/letsencrypt/live/lecture.hduhelp.com/privkey.pem;


    location / {
        proxy_set_header X-Forwarded-For $remote_addr; #主要是要设置这个请求头，让 gin 能够获取到准确客户端的 ip
        proxy_pass http://127.0.0.1:8000;
    }
}

```
