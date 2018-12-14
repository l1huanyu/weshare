package gateway

import (
	"sync"
)

type safaMap struct {
	sync.RWMutex
	Map map[string]func(userID, content string) string
}

func newSafaMap() *safaMap {
	sm := new(safaMap)
	sm.Map = make(map[string]func(userID, content string) string)
	return sm
}

func (sm *safaMap) read(key string) (func(userID, content string) string, bool) {
	sm.RLock()
	value, ok := sm.Map[key]
	sm.RUnlock()
	return value, ok
}

func (sm *safaMap) write(key string, value func(userID, content string) string) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}

var todos *safaMap

func init() {
	todos = newSafaMap()
}

//操作类型
const (
	_GET = "0"
	_SET = "1"
)

const (
	_GET_ACTIVATED = "接下来，选择什么类型呢？\n0：小说\n1：电影"
)

const (
	_SET_ACTIVATED = "接下来，选择什么类型呢？\n0：小说\n1：电影"
)

//Route 接受来自于wxadp层的用户消息
func Route(userID, content string) string {
	if len(userID) == 0 || len(content) == 0 {
		return "错误の输入"
	}

	if todo, ok := todos.read(userID); ok {
		return todo(userID, content)
	}

	return active(userID, content)
}

func active(userID, content string) string {
	switch content {
	case _GET:
		todos.write(userID, getSelecteType)
		return _GET_ACTIVATED
	case _SET:
		todos.write(userID, setSelecteType)
		return _SET_ACTIVATED
	default:
		return "错误の操作类型"
	}
}

func getSelecteType(userID, content string) string {
	//TODO:具体实现，调用下层函数
	todos.write(userID, getNextHop)
	return "getSelectType"
}

func setSelecteType(userID, content string) string {
	//TODO:具体实现，调用下层函数

	return "setSelectType"
}

func getNextHop(userID, content string) string {
	//TODO:具体实现，调用下层函数
	return "选择吧！\n0：被人安利\n1：安利别人"
}
