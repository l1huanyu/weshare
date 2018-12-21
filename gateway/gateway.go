package gateway

import (
	"fmt"
	"weshare/dao"
)

var todos *safeMap

func init() {
	todos = newSafaMap()
}

//操作类型
const (
	_GET = "0"
	_SET = "1"
)

//提示消息
const (
	_Prologue        = "Hi~ o(*￣▽￣*)ブ欢迎光临安利小卖部！\n小店共有安利%d枚，你已寄售安利%d枚~\n\n客官是准备？\n0：被人安利\n1：安利别人"
	_NOT_SUPORT      = "请输入有效数字≧ ﹏ ≦"
	_INTERNAL_ERROR  = "店小二失踪了，再试一次8~"
	_NOT_FOUND       = "客官...8好意思...缺货了...请重新选择类型...(✿◡‿◡)"
	_GET_ACTIVATED   = "接下来，选择什么类型呢？\n\n0：随便来点啥\n1：电影\nn2：游戏\n3：动漫\n4：小说\n5：漫画\n6：音乐\n7：Others~"
	_GET_SELECTED    = "Buling Buling~久等啦~您要的安利已上菜~\n\n%s\n\n0：下一个 1：返回"
	_SET_ACTIVATED   = "接下来，选择什么类型呢？\n\n1：电影\nn2：游戏\n3：动漫\n4：小说\n5：漫画\n6：音乐\n7：Others~"
	_SET_TYPE        = "客官想安利什么呢？"
	_SET_NAME        = "那么在哪里可以买得到呢~请输入获取方式~\nTips: 输入'0'跳过"
	_SET_SOURCE      = "请输入安利理由~\nTips: 输入'0'跳过"
	_SET_DESCRIPTION = "寄售安利成功~\n%s"
)

//Route 接受来自于wxadp层的用户消息
func Route(userID string, content string) string {
	if todo, ok := todos.read(userID); ok {
		return todo(userID, content)
	}
	return active(userID, content)
}

//Realese 释放取消关注的用户的资源
func Realese(userID string) {
	todos.delete(userID)
}

//Prologue 开场白
func Prologue(userID string) string {
	return fmt.Sprintf(_Prologue, dao.CountPosts(), dao.CountPostsByPublisher(userID))
}

//NotSurport 不支持
func NotSurport() string {
	return _NOT_SUPORT
}

func active(userID string, content string) string {
	switch content {
	case _GET:
		todos.write(userID, getSelecteType)
		return _GET_ACTIVATED
	case _SET:
		todos.write(userID, setSelectType)
		return _SET_ACTIVATED
	default:
		return Prologue(userID)
	}
}
