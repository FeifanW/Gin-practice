package api

import (
	"Gin-practice/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UserController struct {
}
type DefaultController struct {
}

func (con UserController) Index(c *gin.Context) {
	// 设置sessions
	session := sessions.Default(c)
	session.Set("username", "张三111")
	session.Save() // 设置session的时候必须调用
}
func (con UserController) GetSession(c *gin.Context) {
	// 设置sessions
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200, "username=%v", username)
}

func (con UserController) List(c *gin.Context) {
	c.String(200, "我是一个api接口-userlist，单独放在api文件夹下")
}

/*
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face")
	// file.Filename 获取文件名称  aaa.png   ./static/upload/aaa.jpg 保存路径
	dst := path.Join("/static/upload", file.Filename)
	if err == nil {
		c.SaveUploadedFile(file, dst)
	}
	//c.String(200, "执行上传")
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
		"dst":      dst,
	})
}

*/

func (con UserController) DoEdit(c *gin.Context) {
	username := c.PostForm("username")
	file1, err1 := c.FormFile("face1")
	// file.Filename 获取文件名称  aaa.png   ./static/upload/aaa.jpg 保存路径
	if err1 == nil {
		dst := path.Join("/static/upload", file1.Filename)
		c.SaveUploadedFile(file1, dst)
	}

	file2, err2 := c.FormFile("face2")
	// file.Filename 获取文件名称  aaa.png   ./static/upload/aaa.jpg 保存路径
	if err2 == nil {
		dst := path.Join("/static/upload", file2.Filename)
		c.SaveUploadedFile(file2, dst)
	}

	//c.String(200, "执行上传")
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}

/*
1.获取上传的文件
2.获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
3.创建图片保存目录 static/upload/20230623
4.生成文件名称和文件保存的目录
5.执行上传
*/
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	// 1.获取上传的文件
	file, err := c.FormFile("face")

	if err == nil {
		// 2.获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			c.String(200, "上传的文件类型不合法")
			return
		}
		// 3.创建图片保存目录 static/upload/20230419
		day := models.GetDay()
		dir := "./static/upload/" + day
		os.MkdirAll(dir, 0666) // 创建文件
		if err != nil {
			fmt.Println(err)
			c.String(200, "MkdirAll失败")
			return
		}
		// 4.生成文件名称和文件保存目录
		fileName := strconv.FormatInt(models.GetUnix(), 10) + extName
		// 5.执行上传
		dst := path.Join(dir, fileName)
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}

func (con DefaultController) Index(c *gin.Context) {
	// 设置cookie
	c.SetCookie("username", "张三", 3600, "/", "localhost", false, false)
}

func (con DefaultController) News(c *gin.Context) {
	// 获取cookie
	username, _ := c.Cookie("username")
	c.String(200, "cookie"+username)
}

// 删除cookie
func (con DefaultController) DeleteCookie(c *gin.Context) {
	// 删除cookie
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(200, "删除成功")
}
