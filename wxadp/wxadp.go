package wxadp

import (
	"net/http"
	"time"
	"weshare/gateway"

	"github.com/labstack/echo"
)

const (
	_APP_ID      = "wx9a5d263b039e6755"
	_SECRET      = "29789cc7663433e29e25ce0697e44aa5"
	_TOKEN       = "l1huanyu"
	_WECHAT_ID   = "hyhappyhouse"
	_EVENT       = "event"
	_TEXT        = "text"
	_SUBSCRIBE   = "subscribe"
	_UNSUBSCRIBE = "unsubscribe"
)

type (
	TextMsgRx struct {
		ToUserName   string `xml:"ToUserName"`   //开发者微信号
		FromUserName string `xml:"FromUserName"` //发送方账号（一个OpenID）
		CreateTime   int    `xml:"CreateTime"`   //消息创建时间
		MsgType      string `xml:"MsgType"`      //text
		Content      string `xml:"Content"`      //文本消息内容
		MsgId        int64  `xml:"MsgId"`        //消息id
		Event        string `xml:"Event"`        //事件类型
	}

	TextMsgTx struct {
		ToUserName   string //接收方账号（收到的OpenID）
		FromUserName string //开发者微信号
		CreateTime   int    //消息创建时间
		MsgType      string //text
		Content      string //回复的消息内容（可换行）
	}
)

func ResponseWechat(c echo.Context) error {
	echostr := c.QueryParam("echostr")
	return c.String(http.StatusOK, echostr)
}

func ReceiveMessage(c echo.Context) error {
	contentTx := ""
	msgRx := new(TextMsgRx)
	err := c.Bind(msgRx)
	if err != nil {
		return err
	}

	if len(msgRx.ToUserName) == 0 || len(msgRx.FromUserName) == 0 || len(msgRx.MsgType) == 0 || len(msgRx.Content) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	switch msgRx.MsgType {
	case _TEXT:
		contentTx = gateway.Route(msgRx.FromUserName, msgRx.Content)
	case _EVENT:
		if msgRx.Event == _SUBSCRIBE {
			contentTx = gateway.Prologue(msgRx.FromUserName)
		} else {
			if msgRx.Event == _UNSUBSCRIBE {
				gateway.Realese(msgRx.FromUserName)
			}
			return c.NoContent(http.StatusOK)
		}
	default:
		contentTx = gateway.NotSurport()
	}

	msgTx := &TextMsgTx{
		ToUserName:   msgRx.FromUserName,
		FromUserName: msgRx.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      _TEXT,
		Content:      contentTx,
	}

	return c.XML(http.StatusOK, msgTx)
}
