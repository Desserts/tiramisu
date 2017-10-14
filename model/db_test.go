package model

import "testing"

var (
	taskModel *TaskStruct
)

func TestInit(t *testing.T) {

}

func TestTaskStruct_AddTask(t *testing.T) {
	url := "https://laily.net"
	types := 1
	err := taskModel.AddTask(url, types)
	if err != nil {
		t.Error(err)
	}
}

func TestTaskStruct_All(t *testing.T) {
	res, err := taskModel.All()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestTaskStruct_GetOneWait(t *testing.T) {
	res, err := taskModel.GetOneWait()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
