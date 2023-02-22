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
