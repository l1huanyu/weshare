package main

import (
	"suren/wxadp"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/wechat", wxadp.ResponseWechat)
	e.POST("/wechat", wxadp.ReceiveMessage)
	e.Start(":8823")
}
