package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.String(200, "hello,world!")
}

func main() {
	// 创建一个默认的路由
	router := gin.Default() // 创建一个路由
	// 绑定路由规则和路由函数，访问/index的路由，将由对应的函数去处理
	router.GET("/index", Index)
	// 启动监听，gin会把web服务运行在本机上0.0.0.0:8080端口上
	router.Run(":8080")
	// 用原生http服务的方式，router.Run的本质就是http.ListenAndServer的进一步封装
	http.ListenAndServe(":8080", router)
}
