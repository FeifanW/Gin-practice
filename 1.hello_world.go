package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default() // 创建一个路由
	router.GET("/index", func(context *gin.Context) {
		context.String(200, "hello,world!")
	})
	router.Run(":8080")
}
