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

##### 查询参数Query：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) // 如果传了多个user，拿到相同的查询参数
    fmt.Println(c.DefaultQuery("addr", "北京")) // 用户没传就使用默认值
}

func main() {
	router := gin.Default()
	router.GET("/query", _query)
	router.Run(":80")
}
```

##### 动态参数Param：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) // 如果传了多个user，拿到相同的查询参数
}

func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id")) // 除了路径之外的信息
	fmt.Println(c.Param("book_id"))
	//http://127.0.0.1:80/param/xxxhhh/bookid13
}

func main() {
	router := gin.Default()
	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.Run(":80")
}
```

##### 表单参数PostForm：

可以接收`mutipart/form-data`和`application/x-www-form-urlencoded`

```go
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "北京")) // 如果用户没有传就使用默认值
	forms, err := c.MultipartForm()              // 接收所有的form参数，包括文件
	fmt.Println(forms, err)
}
func main() {
	router := gin.Default()
	router.POST("/form", _form)
	router.Run(":80")
}
```

**补充：**

###### form-data

就是http请求中的multipart/form-data,它会将表单的数据处理为一条消息，以标签为单元，用分隔符分开。既可以上传键值对，也可以上传文件。当上传的字段是文件时，会有Content-Type来表名文件类型；content-disposition，用来说明字段的一些信息；

由于有boundary隔离，所以multipart/form-data既可以上传文件，也可以上传键值对，它采用了键值对的方式，所以可以上传多个文件。

###### x-www-form-urlencoded

就是application/x-www-from-urlencoded,会将表单内的数据转换为键值对，比如,name=java&age = 23

###### 区别

multipart/form-data：既可以上传文件等二进制数据，也可以上传表单键值对，只是最后会转化为一条信息； x-www-form-urlencoded：只能上传键值对，并且键值对都是间隔分开的。

##### 原始参数GetRawData：

```go
func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	fmt.Println(contentType)
	switch contentType {
	case "application/json": // 解析json数据
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func _raw(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	err := _bindJson(c, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
func main() {
	router := gin.Default()
	router.POST("/raw", _raw)
	router.Run(":80")
}
```

##### 四大请求方式：

Restful风格指的是网络应用中资源定位和资源操作的风格，不是标准也不是协议

GET：从服务器取出资源（一项或多项）

POST：在服务器新建一个资源

PUT：在服务器更新资源（客户端提供完整资源数据）

PATCH：在服务器更新资源（客户端提供需要修改的资源数据）

DELETE：从服务器删除资源

```go
/* 
以文字资源为例
GET    /articles       文章列表
GET    /articles/:id   文章详情
POST   /articles       添加文章
PUT    /articles/:id   修改某一篇文章
DELETE /articles/:id   删除某一篇文章
*/
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	//fmt.Println(contentType)
	switch contentType {
	case "application/json": // 解析json数据
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

// 文章列表页面
func _getList(c *gin.Context) {
	// 搜索和分页
	articleList := []ArticleModel{
		{"Go语言入门", "这篇文章是Go语言入门"},
		{"红宝书", "这篇文章是JS红宝书"},
	}
	c.JSON(200, articleList)
}

func _getDetail(c *gin.Context) {
	// 获取param中的id
	fmt.Println(c.Param("id"))
	article := ArticleModel{"红宝书", "这篇文章是JS红宝书"}
	c.JSON(200, article)
}

// 创建文章
func _create(c *gin.Context) {
	// 接收前端传来的JSON数据
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, article)
}

func _update(c *gin.Context) {
	// 接收前端传来的JSON数据
	fmt.Println(c.Param("id"))
	var article ArticleModel
	err := _bindJson(c, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, article)
}

func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, gin.H{})
}

func main() {
	router := gin.Default()
	//gin.SetMode(gin.DebugMode)
	//router.GET("/articles", _getList)
	//router.GET("/articles", _getDetail)
	//router.POST("/articles", _create)
	//router.POST("/articles/:id", _update)
	router.POST("/articles/:id", _delete)
	router.Run(":80")
}
```

##### 请求头相关：

































