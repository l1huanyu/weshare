package gateway

import "fmt"

const (
	_GET_TYPE_RANDOM = iota
	_GET_TYPE_MOVIE
	_GET_TYPE_NOVEL
	_GET_TYPE_GAME
	_GET_TYPE_ANIMATION
	_GET_TYPE_TELEPLAY
	_GET_TYPE_OTHERS
)

const (
	_GET_NEXT_ONE = iota
	_GET_BACK
)

func getSelecteType(userID string, content int) string {
	if content < 0 || content > _GET_TYPE_OTHERS {
		return _NOT_SURPORT
	}
	//TODO:具体实现，调用controller层函数
	response := ""
	todos.write(userID, getNextOne)
	return fmt.Sprintf(_GET_SELECTED, response)
}

func getNextOne(userID string, content int) string {
	switch content {
	case _GET_NEXT_ONE:
		//TODO:具体实现，调用controller层函数
		response := ""
		todos.write(userID, getNextOne)
		return fmt.Sprintf(_GET_SELECTED, response)
	case _GET_BACK:
		todos.delete(userID)
		return _Prologue
	default:
		return _NOT_SURPORT
	}
}
