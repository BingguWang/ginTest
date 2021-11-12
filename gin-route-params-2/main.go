package main

//当请求带参数时,与结构体相关的参数
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/singer/:id", func(c *gin.Context) {
		s := Singer{}
		if c.ShouldBind(&s) == nil { //用ShouldBind接收结构体绑定的形式的参数
			fmt.Println(s.Age)
			fmt.Println(s.Name)
		} else {
			fmt.Println("没进来")
		}
	})
	r.Run(":8088")
}

type Singer struct {
	Id   int    `json:"id"`
	Age  int    `json:"age" form:"age"`
	Name string `json:"name" form:"name"`
}
