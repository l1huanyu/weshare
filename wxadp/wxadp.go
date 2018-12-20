package wxadp

import (
	"net/http"
	"strconv"
	"time"
	"weshare/gateway"

	"github.com/l1huanyu/suren"
	"github.com/labstack/echo"
)

const (
	_APP_ID      = "wx9a5d263b039e6755"
	_SECRET      = "29789cc7663433e29e25ce0697e44aa5"
	_TOKEN       = "l1huanyu"
	_WECHAT_ID   = "hyhappyhouse"
	_EVENT       = "_event"
	_TEXT        = "text"
	_SUBSCRIBE   = "subscribe"
	_UNSUBSCRIBE = "unsubscribe"
)

var s = suren.New(_APP_ID, _SECRET, _TOKEN)

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
	contentTx := ""
	msgRx := new(suren.TextMsgRx)
	err := c.Bind(msgRx)
	if err != nil {
		return err
	}

	if len(msgRx.ToUserName) == 0 || len(msgRx.FromUserName) == 0 || len(msgRx.MsgType) == 0 || len(msgRx.Content) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	contentRx, err := strconv.Atoi(msgRx.Content)
	if err != nil {
		contentTx = gateway.NotSurport()
		goto RESPONSE
	}

	switch msgRx.MsgType {
	case _TEXT:
		contentTx = gateway.Route(msgRx.FromUserName, contentRx)
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

RESPONSE:
	msgTx := &suren.TextMsgTx{
		ToUserName:   msgRx.FromUserName,
		FromUserName: msgRx.ToUserName,
		CreateTime:   int(time.Now().Unix()),
		MsgType:      _TEXT,
		Content:      contentTx,
	}

	return c.XML(http.StatusOK, msgTx)
}
