package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//step1 获取 http_proxy=192.168.1.2:20080 go get github.com/gin-gonic/gin
//http_proxy=192.168.1.2:20080 这部分是代理设置
func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
	}
	fmt.Println("hello world")
	r.Run(":8080")
}

func GetUsers(c *gin.Context) {
	c.JSON(200, "helloman")
}
