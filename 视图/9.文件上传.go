package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		readerFile, _ := file.Open()
		// 也可以使用os.Create
		writeFile, _ := os.Create("./upload/13.png")
		defer writeFile.Close()
		n, _ := io.Copy(writeFile, readerFile) // 拷贝
		fmt.Println("n=", n)

		//fmt.Println(string(data))
		//fmt.Println(file.Filename)
		fmt.Println(file.Size / 1024) // 单位是字节
		c.SaveUploadedFile(file, "./upload/12.png")
		c.JSON(200, gin.H{"msg": "上传成功"})
	})

	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files, _ := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传%d个文件", len(files))})
	})

	router.Run(":80")
}
