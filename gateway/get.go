package gateway

import (
	"fmt"
	"strconv"
	"weshare/dao"

	"github.com/jinzhu/gorm"
)

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

func getSelecteType(userID string, oldContent string) string {
	content, err := strconv.Atoi(oldContent)
	if err != nil {
		return _NOT_SUPORT
	}
	if content < 0 || content > _GET_TYPE_OTHERS {
		return _NOT_SUPORT
	}
	response := new(dao.Post)
	if content == 0 {
		response, err = dao.QueryPostRandomly()
	} else {
		response, err = dao.QueryPostByType(content)
	}
	if err != nil {
		resp := _INTERNAL_ERROR
		if err == gorm.ErrRecordNotFound {
			resp = _NOT_FOUND
		}
		return resp
	}
	todos.writeState(userID, content)
	todos.write(userID, getNextOne)
	return fmt.Sprintf(_GET_SELECTED, response.Display())
}

func getNextOne(userID string, oldContent string) string {
	content, err := strconv.Atoi(oldContent)
	if err != nil {
		return _NOT_SUPORT
	}
	switch content {
	case _GET_NEXT_ONE:
		//TODO:具体实现，调用controller层函数
		t, _ := todos.readState(userID)
		response := new(dao.Post)
		var err error
		if t == 0 {
			response, err = dao.QueryPostRandomly()
		} else {
			response, err = dao.QueryPostByType(t)
		}
		if err != nil {
			resp := _INTERNAL_ERROR
			if err == gorm.ErrRecordNotFound {
				resp = _NOT_FOUND
			}
			return resp
		}
		todos.write(userID, getNextOne)
		return fmt.Sprintf(_GET_SELECTED, response)
	case _GET_BACK:
		todos.deleteState(userID)
		todos.delete(userID)
		return _Prologue
	default:
		return _NOT_SUPORT
	}
}
