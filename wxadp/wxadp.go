package wxadp

import (
	"net/http"
	"suren/gateway"
	"time"

	"github.com/l1huanyu/suren"
	"github.com/labstack/echo"
)

const (
	APP_ID    = "wx9a5d263b039e6755"
	SECRET    = "29789cc7663433e29e25ce0697e44aa5"
	TOKEN     = "l1huanyu"
	WECHAT_ID = "hyhappyhouse"
	EVENT     = "event"
	TEXT      = "text"
	SUBSCRIBE = "subscribe"
)

type MsgRx struct {
	ToUserName   string `xml:"ToUserName"`   //开发者微信号
	FromUserName string `xml:"FromUserName"` //发送方账号（一个OpenID）
	CreateTime   int    `xml:"CreateTime"`   //消息创建时间
	MsgType      string `xml:"MsgType"`      //text
	Content      string `xml:"Content"`      //文本消息内容
	MsgId        int64  `xml:"MsgId"`        //消息id
	Event        string `xml:"Event"`        //事件类型
}

var s = suren.New(APP_ID, SECRET, TOKEN)

func ResponseWechat(c echo.Context) error {
	echostr := c.QueryParam("echostr")
	if ok, err := s.CheckSignature(&suren.Signature{
		Signature: c.QueryParam("signature"),
		Timestamp: c.QueryParam("timestamp"),
		Nonce:     c.QueryParam("nonce"),
		Echostr:   echostr,
	}); ok && err != nil {
		return c.String(http.StatusOK, echostr)
	}
	return c.NoContent(http.StatusAccepted)
}

func ReceiveMessage(c echo.Context) error {
	content := ""
	msgRx := new(MsgRx)
	err := c.Bind(msgRx)
	if err != nil {
		return err
	}

	if len(msgRx.ToUserName) == 0 || len(msgRx.FromUserName) == 0 || len(msgRx.MsgType) == 0 || len(msgRx.Content) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	switch msgRx.MsgType {
	case TEXT:
		content = gateway.Route(msgRx.FromUserName, msgRx.Content)
	case EVENT:
		if msgRx.Event == SUBSCRIBE {
			content = "选择吧！\n0：被人安利\n1：安利别人"
		} else {
			return c.NoContent(http.StatusOK)
		}
	default:
		content = "不支持の消息类型"
	}

	msgTx := &suren.TextMsgTx{
		ToUserName:   msgRx.FromUserName,
		FromUserName: msgRx.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      TEXT,
		Content:      content,
	}

	return c.XML(http.StatusOK, msgTx)
}
