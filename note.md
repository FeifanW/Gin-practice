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

请求头参数获取

GetHeader可以大小写部分，且返回切片中的第一个数据

```go
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

	router.Run(":80")
}
```

##### 响应头相关：

```go
// 设置响应头
router.GET("/res", func(c *gin.Context) {
    c.Header("Token", "123asdasda45646asd")
    c.Header("Content-Type", "application/text; charset=utf-8") // 浏览器会当成文本直接下载
    c.JSON(0, gin.H{"data": "看看响应头"})
})
```

##### 参数绑定：

gin中的bind可以很方便的进行参数绑定及参数校验

使用这个功能的时候，需要给结构体加上Tag：json 、form、url、xml、yaml

- Must Bind

  不用，校验失败会改状态码

- Should Bind

  可以绑定json、query、param、yaml、xml

  如果校验不通过会返回错误
  
  ```go
  package main
  
  import "github.com/gin-gonic/gin"
  
  type UserInfo struct {
  	Name string `json:"name"`
  	Age  int    `json:"age"`
  	Sex  string `json:"sex"`
  }
  
  func main() {
  	router := gin.Default()
  	router.POST("/", func(c *gin.Context) {
  
  		var userInfo UserInfo
  		err := c.ShouldBind(&userInfo)
  		if err != nil {
  			c.JSON(200, gin.H{"msg": "你错了"})
  			return
  		}
  		c.JSON(200, userInfo)
  	})
  	router.Run(":80")
  }
  ```
  
- ShouldBindQuery

  绑定查询参数：

  tag对应为form

  ```go
  type UserInfo struct {
  	Name string `json:"name" form:"name"`
  	Age  int    `json:"age" form:"age"`
  	Sex  string `json:"sex" form:"sex"`
  }
  
  func main() {
  	router := gin.Default()
  	router.POST("/query", func(c *gin.Context) {
  		var userInfo UserInfo
  		err := c.ShouldBindQuery(&userInfo)
  		if err != nil {
  			fmt.Println(err)
  			c.JSON(200, gin.H{"msg": "你错了"})
  			return
  		}
  		c.JSON(200, userInfo)
  	})
  
  	router.Run(":80")
  }
  ```

- ShouldBindUri

  tag对应uri

  ```go
  type UserInfo struct {
  	Name string `json:"name" form:"name" uri:"name"`
  	Age  int    `json:"age" form:"age" uri:"age"`
  	Sex  string `json:"sex" form:"sex" uri:"sex"`
  }
  
  func main() {
  	router := gin.Default()
  	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {
  		var userInfo UserInfo
  		err := c.ShouldBindUri(&userInfo)
  		if err != nil {
  			fmt.Println(err)
  			c.JSON(200, gin.H{"msg": "你错了"})
  			return
  		}
  		c.JSON(200, userInfo)
  	})
  
  	router.Run(":80")
  }
  ```

- 绑定formData

  会根据请求头中的content-type去自动绑定

  form-data的参数也用这个，tag用form

  默认的tag就是form

  绑定form-data、x-www-form-urlencode

  ```go
  type UserInfo struct {
  	Name string `form:"name"`
  	Age  int    `form:"age"`
  	Sex  string `form:"sex"`
  }
  
  func main() {
  	router := gin.Default()
  	router.POST("/form", func(c *gin.Context) {
  		var userInfo UserInfo
  		err := c.ShouldBind(&userInfo)
  		if err != nil {
  			fmt.Println(err)
  			c.JSON(200, gin.H{"msg": "你错了"})
  			return
  		}
  		c.JSON(200, userInfo)
  	})
  
  	router.Run(":80")
  }
  ```

#### 四、验证器

需要使用参数验证功能，需要加binding tag

##### 常用验证器：

| 配置     | 作用               | 示例                                                |
| -------- | ------------------ | --------------------------------------------------- |
| required | 必填字段           | binding:"min=5"                                     |
| min      | 最小长度           | binding:"max=10"                                    |
| max      | 最大长度           | binding:"eq=3"                                      |
| len      | 长度               | binding:"len=6"                                     |
| eq       | 等于               | binding:"eq=3"                                      |
| ne       | 不等于             | binding:"ne=12"                                     |
| gt       | 大于               | binding:"gt=10"                                     |
| gte      | 大于等于           | binding:"gte=10"                                    |
| lt       | 小于               | binding:"lt=10"                                     |
| lte      | 小于等于           | binding:"lte=10"                                    |
| eqfield  | 等于其他字段的值   | PassWord string `binding:"eqfield=ConfirmPassword"` |
| nefield  | 不等于其他字段的值 |                                                     |
| -        | 忽略字段           | binding:"-"                                         |

```go
package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	//Name       string `json:"name" binding:"required"` // 用户名
	Name     string `json:"name" binding:"min=4,max=6"` // 用户名
	Age      int    `json:"age"`                        // 年龄
	Password string `json:"password"`                   // 密码
	//RePassword string `json:"re_password"`                // 确认密码
	RePassword string `json:"re_password" binding:"eqfield=Password"` // 确认密码
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
		}
		c.JSON(200, gin.H{"msg": user})
	})
	router.Run(":80")
}
```

##### gin内置验证器：

url是uri的子集

```go
/*
枚举  只能是red或green
oneof=red green

字符串
contains=hello   包含hello的字符串
excludes  不包含
startswith 字符串前缀
endswith  字符串后缀

数组
dive  dive后面的验证就是针对数组中的每一个元素

网络验证
ip
ipv4
ipv6
uri     在于I（Identifier）是统一资源标示符，可以唯一标识一个资源
url     在于Locater,是统一资源定位符，提供找到该资源的确切路径

日期验证
datetime=2006-01-02 15:04:05

*/
```

```go
package main

import "github.com/gin-gonic/gin"

type SignUserInfo struct {
	//Name       string `json:"name" binding:"required"` // 用户名
	//Name     string `json:"name" binding:"min=4,max=6"` // 用户名
	//Name string `json:"name" binding:"excludes=f"` // 用户名   不包含f
	//Name string `json:"name" binding:"startswith=f"` // 用户名   前缀为f
	Name     string   `json:"name" binding:"endswith=f"` // 用户名   后缀为f
	LikeList []string `json:"like_list" binding:"required,dive,startswith=like"`
	IP       string   `json:"ip" binding:"ip"`                             // 需要满足ip要求
	Url      string   `json:"url" binding:"url"`                           // 满足url要求
	Uri      string   `json:"uri" binding:"uri"`                           // 满足url要求
	Age      int      `json:"age"`                                         // 年龄
	Password string   `json:"password"`                                    // 密码
	Date     string   `json:"date" binding:"datetime=2006-01-02 15:04:05"` // 1月2日下午3点4分5秒在2006年
	//RePassword string `json:"re_password"`                // 确认密码
	RePassword string `json:"re_password" binding:"eqfield=Password"` // 确认密码
	Sex        string `json:"sex" binding:"oneof=man woman"`          //枚举
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
		}
		c.JSON(200, gin.H{"msg": user})
	})
	router.Run(":80")
}
```

##### 自定义验证的错误信息：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 获取结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj) // 拿到值
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段获取结构体具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required" msg:"用户名校验失败"`
			Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
		return
	})
	router.Run(":80")
}
```

##### 自定义绑定器：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 获取结构体中的msg参数
func _GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj) // 拿到值
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段获取结构体具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}

type User struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"`
	Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
}

func signValid(fl validator.FieldLevel) bool {
	var nameList []string = []string{"辉哥", "龙哥", "灿哥"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.POST("/", func(c *gin.Context) {

		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
		return
	})
	router.Run(":80")
}
```

#### 五、文件

##### 文件上传：

###### 单文件：

SaveUploadFile

```go
c.SaveUploadFile(file,dst)   // 文件对象 文件路径，注意要从新项目根路径开始写
```

```go
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
	router.Run(":80")
}
```

###### 多个文件：

```go
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
```

##### 文件下载：

直接响应一个路径下的文件

```go
c.File("uploads/12.png")
```

有些响应，比如图片，浏览器就会显示这个图片，而不是下载，所以我们需要使浏览器唤起下载行为

```go
c.Header("Content-Type","application/octet-stream")  // 表示是文件流，唤起浏览器下载，一般设置了这个就需要设置文件名
c.Header("Content-Dispositon","attachment;filename="+"test.png")  // 用来指定下载下来的文件名
c.Header("Content-Transfer-Encoding","binary")  // 表示传输过程中的编码形式，乱码问题可能就是因为它
c.File("uploads/test.png")
```

注意：文件下载浏览器可能会有缓存，这个需要注意一下，解决办法就是加查询参数

前端在写的时候要

```js
this.$http({
    method:"post",
    utl:"file/upload",
    data:postData,
    responseType:"blob"   // 前端要加这个
}).then(res=>{})
```

```go
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
```

#### 六、gin中间件和路由

Gin框架允许开发者在处理请求中，加入用户自己的钩子（Hook）函数，这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限教研、数据分页、记录日志、耗时统计等。比如，如果访问一个网页的话，不管访问什么路径都需要进行登录，此时就需要为所有路径处理函数进行统一一个中间件

Gin的中间件必须是一个gin.HandlerFunc类型

##### 单独注册中间件

c.Next()之前的就是请求中间件，之后的就是响应中间件

如果其中一个中间件响应了c.Abort()，后续中间件将不再执行，直接按照顺序走完所有的响应中间件

![image-20230222235317425](D:\practice Space\Gin-practice\assets\image-20230222235317425.png)

```go
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
```

##### 全局中间件和中间件传参：

###### 全局注册中间件：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m10(c *gin.Context) { // 全局中间件
	fmt.Println("m10...in")
	c.JSON(200, gin.H{"msg": "响应被我吃掉了"})
	c.Abort()
	c.Next()
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
		c.Next()
		fmt.Println("index...out")
	})
	router.GET("/m11", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "m11"})
	})
	router.Run(":8080")
}
```

###### 中间件传递数据：

```go
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
```

##### 路由分组：

将一系列路由放到一个组下，统一管理

```go
package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ArticleInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserList(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"book", 21},
		{"li", 22},
		{"wang", 23},
	}
	c.JSON(200, Response{
		Code: 0,
		Data: userList,
		Msg:  "请求成功",
	})
}
func articleList(c *gin.Context) {
	var userList = []ArticleInfo{
		{"go", "从0到1"},
		{"python", "从0到1"},
	}
	c.JSON(200, Response{
		Code: 0,
		Data: userList,
		Msg:  "请求成功",
	})
}

func main() {
	router := gin.Default()

	api := router.Group("api")
	userManger := api.Group("user_manager")
	{
		userManger.GET("/users", UserList) // 访问需要/api/user_manager/users
	}
	articleManager := api.Group("article_manager")
	{
		articleManager.GET("/articles", articleList)
	}
	//api.GET("/users", UserList)        // 访问需要/api/users
	//router.GET("/users", UserList)
	router.Run(":8080")
}
```

##### 路由分组中间件：

```go

```































