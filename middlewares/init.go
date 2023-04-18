package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	// 判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	cp := c.Copy()
	// 定义一个goroutine统计日志
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done in path" + cp.Request.URL.Path)
	}()
}
