package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf(
			"%s %s %s %d \n",
			httpMethod,
			absolutePath,
			handlerName,
			nuHandlers,
		)
	}
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"[FENG]%s |%s %d %s| %s %s %s  %s\n",
			params.TimeStamp.Format("2006-01-02 15:04:05"),
			//params.StatusCode,
			//params.Method,
			params.StatusCodeColor(), params.StatusCode, params.Method, params.ResetColor(), // 加颜色
			params.MethodColor(), // 可以修改颜色
			params.Path,
		)
	}))
	router.GET("/index", func(c *gin.Context) {})
	router.POST("/users", func(c *gin.Context) {})
	router.POST("/articles", func(c *gin.Context) {})
	router.DELETE("/articles/:id", func(c *gin.Context) {})
	fmt.Println(router.Routes()) // 会把所有路由打印下来
	router.Run(":8080")
}
