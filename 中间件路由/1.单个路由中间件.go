package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//c.Abort() // 使用这个方法之后，就不会响应后面的了
	c.Next()
	fmt.Println("m1 out ...")
	//c.JSON(200, gin.H{"msg": "index"})

}

func index(c *gin.Context) {
	fmt.Println("index ... in")
	c.JSON(200, gin.H{"msg": "index"})
	c.Next()
	fmt.Println("index ...out")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Next()
	//c.JSON(200, gin.H{"msg": "index"})
	fmt.Println("m2 out ...")
}

func main() {
	router := gin.Default()
	router.GET("/", m1, index, m2)
	router.Run(":8080")
}
