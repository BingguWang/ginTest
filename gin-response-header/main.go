package main

import "github.com/gin-gonic/gin"

//响应头的设置

func main() {
	r := gin.Default()

	//Gin的Header方法可以设置响应头，默认是k-v方式
	r.GET("/login", func(c *gin.Context) {
		c.Header("name", "wbing") //value得是string
		c.Header("status", "200")
		c.Header("token", "SADOIHSABWQJBJ2312IU321H")
	})

	r.Run(":8088")
}
