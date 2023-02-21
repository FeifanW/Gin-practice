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
