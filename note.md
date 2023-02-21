#### 一、环境搭建：

Gin是一个很受欢迎的Go web框架

下载前在Goland最好配置一下代理

```go
GOPROXY=https://goproxy.cn,direct
```

在Goland中创建Go项目之后，下载Gin的依赖

```go
go get -u github.com/gin-gonic/gin
```

安装postman

修改ip为内网ip

```go
router.Run("0.0.0.0:8080")
```

#### 二、响应

##### 响应json：

状态码200等价于http.StatusOk

```go
package main

import "github.com/gin-gonic/gin"

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

func main() {
	router := gin.Default()
	router.GET("/", _string)
	router.GET("/json", _json)
	router.Run(":80")
}
```

##### 响应xml和yaml：

```go
// 响应xml
func _xml(c *gin.Context) {
	c.XML(200, gin.H{"userName": "芒果", "age": "26", "status": http.StatusOK, "data": gin.H{"user": "菠萝"}})
}

// 响应yaml
func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{"userName": "芒果", "age": "26", "status": http.StatusOK, "data": gin.H{"user": "菠萝"}})
}

func main() {
	router := gin.Default()
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.Run(":80")
}
```

##### 响应html：

```go
// 响应html
func _html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"username": "铃兰"})
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*") // 要把html文件加载进来
	router.GET("/html", _html)
	router.Run(":80")
}
```

##### 文件响应：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(200, "你好呀")
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
	router.Run(":80")
}

```

##### 重定向：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	router.GET("/", _string)
	router.GET("/baidu", _redirect)

	router.Run(":80")
}

```

#### 三、请求













