package main

import "github.com/gin-gonic/gin"

//corsHeader 添加响应头，实现跨域使用、同时实现允许使用 REST 中的方法
func corsHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
	return
}
