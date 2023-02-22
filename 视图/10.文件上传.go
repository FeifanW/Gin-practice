package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/download", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")            // 表示是文件流，唤起浏览器下载，一般设置了这个就需要设置文件名
		c.Header("Content-Dispositon", "attachment;filename="+"12.png") // 用来指定下载下来的文件名
		c.File("upload/12.png")                                         // 这样会直接在网页预览
	})
	router.Run(":8080")
}
