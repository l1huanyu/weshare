package gateway

import (
	"fmt"
	"strconv"
	"weshare/dao"
)

func setSelectType(userID string, oldContent string) string {
	content, err := strconv.Atoi(oldContent)
	if err != nil {
		return _NOT_SUPORT
	}
	p := new(dao.Post)
	p.Publisher = userID
	p.Type = content
	p.State = dao.CREATED
	err = p.Create()
	if err != nil {
		return _INTERNAL_ERROR
	}
	todos.writeHopMap(userID, setSelectName)
	return _SET_TYPE
}

func setSelectName(userID string, content string) string {
	p, err := dao.QueryUnfinishedPost()
	if err != nil {
		return _INTERNAL_ERROR
	}
	p.Name = content
	p.State = dao.SET_NAME
	p.Version++
	err = p.Update()
	if err != nil {
		return _INTERNAL_ERROR
	}
	todos.writeHopMap(userID, setSelectSource)
	return _SET_NAME
}

func setSelectSource(userID string, content string) string {
	p, err := dao.QueryUnfinishedPost()
	if err != nil {
		return _INTERNAL_ERROR
	}
	if content != "0" {
		p.Source = content
	}
	p.State = dao.SET_SOURCE
	p.Version++
	err = p.Update()
	if err != nil {
		return _INTERNAL_ERROR
	}
	todos.writeHopMap(userID, setSelectDescription)
	return _SET_SOURCE
}

func setSelectDescription(userID string, content string) string {
	p, err := dao.QueryUnfinishedPost()
	if err != nil {
		return _INTERNAL_ERROR
	}
	if content != "0" {
		p.Description = content
	}
	p.State = dao.SUCCEED
	p.Version++
	err = p.Update()
	if err != nil {
		return _INTERNAL_ERROR
	}
	todos.deleteHopMap(userID)
	return fmt.Sprintf(_SET_DESCRIPTION, Prologue(userID))
}
