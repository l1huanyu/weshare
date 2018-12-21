package main

import (
	"weshare/dao"
	"weshare/wxadp"

	"github.com/labstack/echo"
)

func main() {
	defer func() {
		dao.CloseDB()
	}()
	e := echo.New()
	e.GET("/wechat", wxadp.ResponseWechat)
	e.POST("/wechat", wxadp.ReceiveMessage)
	e.Start(":8823")
}
