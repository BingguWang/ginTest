package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//什么是路由：就是一个http请求找到对应的处理器函数
func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则，绑定执行的函数

	r.GET("/", func(c *gin.Context) { //gin.Context允许我们在中间件之间传递变量,第二个参数处理器函数HanldleFunc，和http包的不一样
		c.String(http.StatusOK, "get hello world\n") //c.String会把传入的内容写入到response body中
		// c.JSON(200, gin.H{
		// 	"message": "hello world",
		// })
	})
	// r.POST("/post", postHandler)
	// r.PUT("/put", func(c *gin.Content) {
	// 	c.String(http.StatusOK, "put hello world")
	// })
	//指定端口号
	r.Run(":8088")
}

func postHandler(c *gin.Context) {
	c.String(http.StatusOK, "post hello world")
}
