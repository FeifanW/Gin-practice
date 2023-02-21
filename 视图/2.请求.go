package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user"))           // 如果传了多个user，拿到相同的查询参数
	fmt.Println(c.DefaultQuery("addr", "北京")) // 用户没传就使用默认值
}

func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id")) // 除了路径之外的信息
	fmt.Println(c.Param("book_id"))
	//http://127.0.0.1:80/param/xxxhhh/bookid13
}

func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "北京")) // 如果用户没有传就使用默认值
	forms, err := c.MultipartForm()                // 接收所有的form参数，包括文件
	fmt.Println(forms, err)
}

//func _bindJson(c *gin.Context, obj any) (err error) {
//	body, _ := c.GetRawData()
//	contentType := c.GetHeader("Content-Type")
//	fmt.Println(contentType)
//	switch contentType {
//	case "application/json": // 解析json数据
//		err := json.Unmarshal(body, &obj)
//		if err != nil {
//			fmt.Println(err.Error())
//			return err
//		}
//	}
//	return nil
//}

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
	router.GET("/query", _query)
	router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.Run(":80")
}
