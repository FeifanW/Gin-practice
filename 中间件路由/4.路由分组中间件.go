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

func MiddleWare(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "1234" {
		c.Next()
		return
	}
	c.JSON(200, Response{
		Code: 200,
		Data: nil,
		Msg:  "权限验证失败",
	})
}

func main() {
	router := gin.Default()
	//router := gin.New()  // 不含任何中间件

	api := router.Group("api")
	userManger := api.Group("user_manager").Use(MiddleWare)
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
