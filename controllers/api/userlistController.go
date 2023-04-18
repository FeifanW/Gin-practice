package api

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (con UserController) List(c *gin.Context) {
	c.String(200, "我是一个api接口-userlist，单独放在api文件夹下")
}
