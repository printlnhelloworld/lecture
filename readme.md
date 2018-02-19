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

使用 dep 做包管理，后台的所有源码都在 `src/backend/` 里面，其中的 `vendor` 是后端依赖目录。不会包含到目录里面

采用 Gin 做框架，同时使用 go-bindata 来包含前端生成文件，达到单文件发行的目的。

## 版本控制

使用 Git 做版本管理

## 怎么编译运行
