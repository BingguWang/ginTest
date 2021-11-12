package main

import (
	"github.com/gin-gonic/gin"
)

//Gin的hTTP请求的响应
//gin.Context上下文支持多种返回处理结果，字符串，JSON，XML等格式
func main() {
	r := gin.Default()

	//字符串响应
	r.GET("/string", func(c *gin.Context) {
		c.String(200, "hello %s ,你好帅", "wbing")
	})

	//JSON，这是最常用的
	r.GET("/json", func(c *gin.Context) {
		s := &Singer{
			Name: "dt",
			Age:  52,
		}
		c.JSON(200, s) //传入结构体的指针,响应结果为{"Name":"dt","Age":52}
	})

	//XML格式
	r.GET("/xml", func(c *gin.Context) {
		s := &Singer{
			Name: "dt",
			Age:  52,
		}
		c.XML(200, s) //xml的根节点的名字默认是结构体的名字
	})

	//文件格式
	r.GET("/file", func(c *gin.Context) {
		// c.File("C:/Users/WBing/Pictures/rmrf.gif") //参数是文件地址
		c.File("./rmrf.gif") //参数是文件地址
	})

	// 4.YAML响应
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan"})
	})

	r.Run(":8088")
}

type Singer struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}
