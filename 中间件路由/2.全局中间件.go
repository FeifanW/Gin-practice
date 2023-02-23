package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

func m10(c *gin.Context) { // 全局中间件
	fmt.Println("m10...in")
	c.JSON(200, gin.H{"msg": "响应被我吃掉了"})
	c.Set("name", "zzz")
	c.Set("user", User{
		Name: "test",
		Age:  25,
	})
	//c.Abort()
	//c.Next()
	fmt.Println("m10...out")
}
func m11(c *gin.Context) { // 全局中间件
	fmt.Println("m11...in")
	c.Next()
	fmt.Println("m11...out")
}

func main() {
	router := gin.Default()
	router.Use(m10, m11) // 注册中间件
	router.GET("/m10", func(c *gin.Context) {
		fmt.Println("index...in")
		c.JSON(200, gin.H{"msg": "m10"})
		name, _ := c.Get("name")
		fmt.Println(name)
		_user, _ := c.Get("user")
		user := _user.(User) // 如果需要用到user中的某一个字段，需要断言
		fmt.Println(user.Name)
		c.Next()
		fmt.Println("index...out")
	})
	router.GET("/m11", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "m11"})
	})
	router.Run(":8080")
}
