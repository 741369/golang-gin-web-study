package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
	//这里我们导入已经集成的 mysql 驱动，当然也可以导入原版的 import _ "github.com/go-sql-driver/mysql" 一样的
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	//"os"
)

//step1 获取 http_proxy=192.168.1.2:20080 go get github.com/gin-gonic/gin
//http_proxy=192.168.1.2:20080 这部分是代理设置
//
//step2 http_proxy=192.168.1.2:20080 go get -u github.com/jinzhu/gorm
//
//step3 http_proxy=192.168.1.2:20080 go get github.com/go-sql-driver/mysql
//
//step4 http_proxy=192.168.1.2:20080 go get github.com/sirupsen/logrus
//
//	step4.1 http_proxy=192.168.1.2:20080 go get github.com/rifflock/lfshook 写入到文件吧大兄弟
type User struct {
	ID        uint `gorm:"primary_key`
	Uname     string
	CreatedAt time.Time
}

//step4.1
var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./info.log",
		logrus.ErrorLevel: "./info.log",
	}
	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log
}

func init() {
	Log = NewLogger()
	// Log as JSON instead of the default ASCII formatter.
	//Log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//Log.SetLevel(log.WarnLevel)
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

	//step4
	// log.Debug("Useful debugging information.")
	// log.Info("Something noteworthy happened!")
	// log.Warn("You should probably take a look at this.")
	// log.Error("Something failed but I'm not quitting.")

	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// 	"size":   10,
	// }).Warn("A group of walrus emerges from the ocean")

	Log.Warn("You should probably take a look at this.")
	Log.Error("Something failed but I'm not quitting.")
	Log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Error("A group of walrus emerges from the ocean")
	r.Run(":8080")

}

func GetUsers(c *gin.Context) {
	c.JSON(200, "hello world")
}
