package main

import (
	"Gin-practice/controllers/api"
	"Gin-practice/middlewares"
	"Gin-practice/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

type H map[string]interface{}

//type Article struct {
//	Title   string `json:"title"`
//	Desc    string `json:"desc"`
//	Content string `json:"content"`
//}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-01 15:04:05")
}

type Userinfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"password"`
}

type Article struct {
	Title   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}

func initMiddlewareOne(c *gin.Context) {
	fmt.Println("我是中间件One")
	// 调用该请求的剩余处理程序
	c.Set("TEST", "测试在中间件传值")
	c.Next()
	fmt.Println("我是一个中间件One")
}

func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("我是中间件Two")
	test, _ := c.Get("TEST")
	fmt.Println("中间件传值获取的信息", test)
	// 调用该请求的剩余处理程序
	c.Next()
	fmt.Println("我是一个中间件Two")
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 自定义模板函数，注意这个函数要放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	// 配置模板的文件，多层级的时候就要这样配置
	r.LoadHTMLGlob("templates/**/*")
	// 配置静态服务web目录  第一个参数表示路由，第二个参数表示映射目录
	r.Static("/static", "././static")

	// 全局中间件
	r.Use(initMiddlewareOne, initMiddlewareTwo)

	//演示中间件
	r.GET("/", func(c *gin.Context) {
		fmt.Println("这是一个首页")
		c.String(200, "gin首页")
	})

	// 配置路由
	/*
		r.GET("/", func(c *gin.Context) {
			username := c.Query("username")
			age := c.Query("age")
			page := c.DefaultQuery("page", "1")
			c.JSON(http.StatusOK, gin.H{
				"username": username,
				"age":      age,
				"page":     page,
			})
			//c.String(200, "值:%v", "你好gin")
		})
		r.GET("/news", func(c *gin.Context) {
			c.String(200, "我是新闻页面")
		})
		r.GET("/json", func(c *gin.Context) {
			c.JSON(200, map[string]interface{}{
				"success": true,
				"msg":     "你好gin",
			})
		})
		r.GET("/ginH", func(c *gin.Context) {
			c.JSON(200, gin.H{ // 可以这样简写
				"success": true,
				"msg":     "你好gin",
			})
		})

		// 获取GET POST 传递的数据绑定到结构体
		r.GET("/getUser", func(c *gin.Context) {
			user := &Userinfo{}
			if err := c.ShouldBind(&user); err == nil {
				fmt.Printf("%#v", user)
				c.JSON(http.StatusOK, user)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			}
		})

		//r.GET("/json2", func(c *gin.Context) {
		//	a := &Article{
		//		Title:   "我是一个标题",
		//		Desc:    "描述",
		//		Content: "测试内容",
		//	}
		//	c.JSON(200, a)
		//})
		//// jsonp和json的区别就是可以执行回调函数
		//r.GET("/jsonp", func(c *gin.Context) {
		//	a := &Article{
		//		Title:   "我是一个jsonp",
		//		Desc:    "描述",
		//		Content: "测试内容",
		//	}
		//	c.JSONP(200, a)
		//})
		// 返回XML数据
		r.POST("/xml", func(c *gin.Context) {
			article := &Article{}
			xmlSliceData, _ := c.GetRawData() // 获取 c.Request.Body 读取请求数据
			if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
				c.JSON(http.StatusOK, article)
				fmt.Printf("%#v", article)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			}
		})

		// 动态路由传值
		r.GET("/list/:cid", func(c *gin.Context) {
			cid := c.Param("cid")
			c.String(200, "%v", cid)
		})

		// 返回html
		//r.GET("/html", func(c *gin.Context) {
		//	a := &Article{
		//		Title:   "我是一个标题",
		//		Desc:    "描述",
		//		Content: "测试内容",
		//	}
		//	c.HTML(http.StatusOK, "test/index.html", gin.H{
		//		"title": "我是后台的数据",
		//		"news":  a,
		//	})
		//})
		r.POST("/add", func(c *gin.Context) {
			username := c.PostForm("username")
			password := c.PostForm("password")
			age := c.DefaultPostForm("age", "20")
			c.JSON(http.StatusOK, gin.H{
				"username": username,
				"password": password,
				"age":      age,
			})
		})

	*/

	//defaultRouters := r.Group("/")
	//{
	//	defaultRouters.GET("/", func(c *gin.Context) {
	//		c.String(200, "首页")
	//	})
	//	defaultRouters.GET("/news", func(c *gin.Context) {
	//		c.String(200, "新闻")
	//	})
	//}

	apiRouters := r.Group("/api", middlewares.InitMiddleware)
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是一个api接口")
		})
		apiRouters.GET("/userlist", api.UserController{}.List)
		r.GET("/plist", func(c *gin.Context) {
			c.String(200, "我是一个api接口-plist")
		})
	}

	r.Run() // 启动web服务
}
