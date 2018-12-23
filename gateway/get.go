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
	_GET_TYPE_TELEPLAY
	_GET_TYPE_GAME
	_GET_TYPE_ANIMATION
	_GET_TYPE_NOVEL
	_GET_TYPE_COMIC
	_GET_TYPE_OTHERS
)

const (
	_GET_LIKE = iota
	_GET_NEXT_ONE
	_GET_BACK
)

func getSelectType(userID string, oldContent string) string {
	content, err := strconv.Atoi(oldContent)
	if err != nil {
		return _NOT_SUPORT
	}
	if content < 0 || content > _GET_TYPE_OTHERS {
		return _NOT_SUPORT
	}
	response := new(dao.Post)
	if content == 0 {
		response, err = dao.QueryPostRandomly(todos.readReadMap(userID))
	} else {
		response, err = dao.QueryPostByType(content, todos.readReadMap(userID))
	}
	if err != nil {
		resp := _INTERNAL_ERROR
		if err == gorm.ErrRecordNotFound {
			resp = _NOT_FOUND
			if content == _GET_TYPE_RANDOM {
				todos.deleteHopMap(userID)
				resp = resp + "\n" + Prologue(userID)
			}
		}
		return resp
	}
	todos.writeTypeMap(userID, content)
	todos.writeReadMap(userID, response.ID)
	todos.writeHopMap(userID, getNextOne)
	return fmt.Sprintf(_GET_SELECTED, response.Display())
}

func getNextOne(userID string, oldContent string) string {
	content, err := strconv.Atoi(oldContent)
	if err != nil {
		return _NOT_SUPORT
	}
	switch content {
	case _GET_LIKE:
		if postID, ok := todos.readCurrentMap(userID); ok {
			dao.Like(userID, postID)
		}
		fallthrough
	case _GET_NEXT_ONE:
		t, _ := todos.readTypeMap(userID)
		response := new(dao.Post)
		var err error
		if t == 0 {
			response, err = dao.QueryPostRandomly(todos.readReadMap(userID))
		} else {
			response, err = dao.QueryPostByType(t, todos.readReadMap(userID))
		}
		if err != nil {
			resp := _INTERNAL_ERROR
			if err == gorm.ErrRecordNotFound {
				resp = _NOT_FOUND
				if t == _GET_TYPE_RANDOM {
					todos.deleteTypeMap(userID)
					todos.deleteHopMap(userID)
					resp = resp + "\n" + Prologue(userID)
				}
			}
			return resp
		}
		todos.writeReadMap(userID, response.ID)
		todos.writeHopMap(userID, getNextOne)
		return fmt.Sprintf(_GET_SELECTED, response.Display())
	case _GET_BACK:
		todos.deleteTypeMap(userID)
		todos.deleteHopMap(userID)
		return Prologue(userID)
	default:
		return _NOT_SUPORT
	}
}
