package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(200, "你好呀")
}

func _json(c *gin.Context) {
	// json响应结构体
	//type Userinfo struct {
	//	UserName string `json:"user_name"`
	//	Age      int    `json:"age"`
	//	Password string `json:"-"` // 忽略会转换成json
	//}
	//user := Userinfo{
	//	UserName: "百香果",
	//	Age:      23,
	//}

	// json响应map
	//userMap := map[string]string{
	//	"username": "草莓",
	//	"age":      "25",
	//}

	// 直接响应json
	c.JSON(200, gin.H{"userName": "芒果", "age": "26"})

	//c.JSON(200, user)
	//c.JSON(200, userMap)
}

// 响应xml
func _xml(c *gin.Context) {
	c.XML(200, gin.H{"userName": "芒果", "age": "26", "status": http.StatusOK, "data": gin.H{"user": "菠萝"}})
}

// 响应yaml
func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{"userName": "芒果", "age": "26", "status": http.StatusOK, "data": gin.H{"user": "菠萝"}})
}

// 响应html
func _html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"username": "铃兰"})
}

// 重定向
func _redirect(c *gin.Context) {
	// 第一个是重定向的状态码
	//c.Redirect(301, "https://www.baidu.com")
	//c.Redirect(302, "https://www.cnblogs.com/oaoa/")
	//也可以直接写路径
	c.Redirect(302, "/html")
}

func main() {
	router := gin.Default()
	// 加载模板目录下的所有模板文件
	router.LoadHTMLGlob("templates/*") // 要把html文件加载进来
	// 在go中，没有相对文件的路径，只有相对项目的路径
	// 这样可以做到只拿到static/static下面的文件，static下面的文件不指定无法获取到
	// 网页请求静态目录的前缀，后面是http.Dir方法，是一个目录，注意前缀不要重复
	router.StaticFS("/static", http.Dir("static/static"))
	// 配置单个文件，网页请求的路由，文件的路径
	router.StaticFile("/cat.png", "static/cat.png")
	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)

	router.Run(":80")
}
