package main

//当请求带参数时,与结构体相关的参数，或者是文件参数
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//与结构体绑定的形式的参数
	r.POST("/singer/:id", func(c *gin.Context) {
		s := Singer{}
		if c.ShouldBind(&s) == nil { //用ShouldBind接收结构体绑定的形式的参数
			fmt.Println(s.Age)
			fmt.Println(s.Name)
		} else {
			fmt.Println("没进来")
		}
	})

	//文件类型参数
	r.POST("/upload", func(c *gin.Context) {
		// file, err := c.FormFile("file") //FormFile方法获取文件参数,但是没有FileHeader头，只有multipart.File
		_, headers, err := c.Request.FormFile("file") //(multipart.File, *multipart.FileHeader, error)
		if err != nil {
			c.String(500, "上传图片出错")
		}
		if headers.Size > 1024*1024*30 { //大于30MB
			c.String(http.StatusForbidden, "文件不能大于30MB")
			return
		}
		fmt.Println(headers.Header)
		if headers.Header.Get("Content-type") != "audio/mpeg" {
			c.String(http.StatusForbidden, "只能上传音频文件")
			return
		}
		c.SaveUploadedFile(headers, "xxxx") //保存文件,第二个参数是保存文件路径
		c.String(http.StatusOK, "xxxx")
	})
	r.Run(":8088")
}

type Singer struct {
	Id   int    `json:"id"`
	Age  int    `json:"age" form:"age"`
	Name string `json:"name" form:"name"`
}
