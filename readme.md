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
