package model

import (
	"errors"
)

const (
	TASK_TABLE_NAME = "task"
)

const (
	TASK_TYPE_NORMAL = iota
	TASK_TYPE_BT
)
const (
	TASK_STATUS_WAIT = iota
	TASK_STATUS_RUNING
	TASK_STATUS_FAIL
	TASK_STATUS_SUCC
)

var (
	TaskTypesContent = []string{
		"普通",
		"BT",
	}
	TaskStatusContent = []string{
		"等待中",
		"下载中",
		"失败",
		"完成",
	}
)

type TaskStruct struct {
	Id     int
	Url    string
	Types  int `db:"type"`
	Status int
}

func (t *TaskStruct) AddTask(url string, types int) error {
	query := "INSERT INTO " + TASK_TABLE_NAME + " (url, `type`, status) VALUES(?,?,?)"
	res, err := db.Exec(query, url, types, TASK_STATUS_WAIT)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 1 {
		return nil
	} else {
		return errors.New("insert error")
	}
}

func (t *TaskStruct) All() ([]TaskStruct, error) {
	var tasks []TaskStruct
	query := "SELECT * FROM " + TASK_TABLE_NAME + " LIMIT 100"
	err := db.Select(&tasks, query)
	return tasks, err
}

func (t *TaskStruct) GetOneWait() (TaskStruct, error) {
	var tt TaskStruct
	query := "SELECT * FROM " + TASK_TABLE_NAME + " WHERE `type` = ? LIMIT 1"
	err := db.Get(&tt, query, TASK_STATUS_WAIT)
	return tt, err
}

func (t *TaskStruct) UpdateGidAndStatus(id int, gid string, status int) error {
	query := "UPDATE " + TASK_TABLE_NAME + " SET gid = ?, status = ? WHERE id = ?"
	res, err := db.Exec(query, gid, status, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
