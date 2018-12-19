package gateway

var todos *safeMap

func init() {
	todos = newSafaMap()
}

//操作类型
const (
	_GET = iota
	_SET
)

//提示消息
const (
	_Prologue      = "选择吧！\n0：被人安利\n1：安利别人"
	_NOT_SURPORT   = "请输入有效数字≧ ﹏ ≦"
	_GET_ACTIVATED = "接下来，选择什么类型呢🧐？\n0：小说 1：电影\n2：随便来点啥"
	_GET_SELECTED  = "Buling Buling~久等啦~您要的安利已上菜~\n%s\n0：下一个 1：返回"
	_SET_ACTIVATED = "接下来，选择什么类型呢🧐？\n0：小说\n1：电影"
)

//Route 接受来自于wxadp层的用户消息
func Route(userID string, content int) string {
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
func Prologue() string {
	return _Prologue
}

//NotSurport 不支持
func NotSurport() string {
	return _NOT_SURPORT
}

func active(userID string, content int) string {
	switch content {
	case _GET:
		todos.write(userID, getSelecteType)
		return _GET_ACTIVATED
	case _SET:
		todos.write(userID, setSelecteType)
		return _SET_ACTIVATED
	default:
		return _Prologue
	}
}
