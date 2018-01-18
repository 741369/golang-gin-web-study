package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
	//这里我们导入已经集成的 mysql 驱动，当然也可以导入原版的 import _ "github.com/go-sql-driver/mysql" 一样的
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//step1 获取 http_proxy=192.168.1.2:20080 go get github.com/gin-gonic/gin
//http_proxy=192.168.1.2:20080 这部分是代理设置
//
//step2 http_proxy=192.168.1.2:20080 go get -u github.com/jinzhu/gorm
//
//step3 http_proxy=192.168.1.2:20080 go get github.com/go-sql-driver/mysql
//
type User struct {
	ID        uint `gorm:"primary_key`
	Uname     string
	CreatedAt time.Time
}

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
	}
	db, err := gorm.Open("mysql", "golang:Password&123@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err == nil {
		fmt.Println("DB connect success")
	}

	user := User{Uname: "test1"}
	fmt.Println(db.NewRecord(user)) // => returns `true` as primary key is blank

	ret := db.Create(&user)
	fmt.Println(user)
	fmt.Println(ret.Error)
	r.Run(":8080")

}

func GetUsers(c *gin.Context) {
	c.JSON(200, "hello world")
}
