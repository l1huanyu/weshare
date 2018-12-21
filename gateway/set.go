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
	todos.write(userID, setSelectName)
	return _SET_NAME
}

func setSelectName(userID string, content string) string {
	p, err := dao.QueryUnfinishedPost()
	if err != nil {
		return _INTERNAL_ERROR
	}
	p.Name = content
	p.State = dao.SET_NAME
	err = p.Update()
	if err != nil {
		return _INTERNAL_ERROR
	}
	todos.write(userID, setSelectSource)
	return _SET_SOURCE
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
	err = p.Update()
	if err != nil {
		return _INTERNAL_ERROR
	}
	todos.write(userID, setSelectDescription)
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
	err = p.Update()
	if err != nil {
		return _INTERNAL_ERROR
	}
	todos.delete(userID)
	return fmt.Sprintf(_SET_DESCRIPTION, Prologue(userID))
}
