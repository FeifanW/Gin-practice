package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type H map[string]interface{}

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-01 15:04:05")
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 自定义模板函数，注意这个函数要放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	// 配置模板的文件，多层级的时候就要这样配置
	r.LoadHTMLGlob("templates/**/*")
	// 配置静态服务web目录  第一个参数表示路由，第二个参数表示映射目录
	r.Static("/static", "././static")
	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "你好gin")
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
	r.GET("/json2", func(c *gin.Context) {
		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})
	// jsonp和json的区别就是可以执行回调函数
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "我是一个jsonp",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})
	// 返回XML数据
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好gin，我是xml",
		})
	})
	// 返回html
	r.GET("/html", func(c *gin.Context) {
		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.HTML(http.StatusOK, "test/index.html", gin.H{
			"title": "我是后台的数据",
			"news":  a,
		})
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "我是post请求返回的数据")
	})
	r.Run() // 启动web服务
}
