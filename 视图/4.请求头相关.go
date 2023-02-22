package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {

	router := gin.Default()
	// 请求头的各种获取方式
	router.GET("/", func(c *gin.Context) {
		//首字母大小写不区分，单词与单词之间用-连接
		// 用于获取一个请求头
		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.GetHeader("User-agent"))
		fmt.Println(c.GetHeader("user-Agent"))

		//Header 是一个普通的map[string] []string
		fmt.Println(c.Request.Header)
		// 如果是使用Get方法或者 .GetHeader, 哪么可以不用区分大小写，并且返回第一个value
		fmt.Println(c.Request.Header.Get("User-Agent"))
		fmt.Println(c.Request.Header["User-Agent"])
		// 如果是map的取值方式，请注意大小写问题
		fmt.Println(c.Request.Header["user-agent"])

		// 自定义的请求头，用Get方法也是免大小写
		fmt.Println(c.Request.Header.Get("Token"))
		fmt.Println(c.Request.Header.Get("token"))
		c.JSON(200, gin.H{"msg": "成功"})
	})

	// 爬虫和用户区别对待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		// 用正则去匹配
		// 字符串的包含匹配
		if strings.Contains(userAgent, "python") {
			// 爬虫来了
			c.JSON(0, gin.H{"data": "这是响应给爬虫的数据"})
			return
		}
	})

	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "123asdasda45646asd")
		c.Header("Content-Type", "application/text; charset=utf-8") // 浏览器会当成文本直接下载
		c.JSON(0, gin.H{"data": "看看响应头"})
	})

	router.Run(":80")
}
