package main

import (
	"flag"
	"fmt"
	"github.com/NoahShen/aria2rpc"
	log "github.com/golang/glog"
	"tiramisu/config"
	"tiramisu/model"
)

var (
	statusMap = map[string]int{
		"active":   model.TASK_STATUS_RUNING,
		"waiting":  model.TASK_STATUS_WAIT,
		"complete": model.TASK_STATUS_SUCC,
		"error":    model.TASK_STATUS_FAIL,
	}
)

type AriaUtil struct{}

func init() {
	flag.Parse()
	flag.Lookup("alsologtostderr").Value.Set("true")
	aria2rpc.RpcUrl = config.RPC_URL
}

func getOneToDownload() bool {
	task, err := taskModel.GetOneWait()
	if err != nil {
		log.Error("get task error: ", err)
	}
	if task.Id == 0 {
		return false
	}
	switch task.Types {
	case model.TASK_TYPE_NORMAL:
		ariaUtil.addUriNormal(task.Id, task.Url)
	}
	return true
}

func (a *AriaUtil) addUriNormal(id int, url string) (string, error) {
	param := make(map[string]interface{})
	gid, err := aria2rpc.AddUri(url, param)
	if err != nil {
		log.Error("add uri error ", err)
		return "", err
	}
	log.Infof("id:%d, gid:%s", id, gid)
	status, statusCode := a.getStatus(gid)
	log.Infof("status info. gid:%s, status:%s", gid, status)
	err = taskModel.UpdateGidAndStatus(id, gid, statusCode)
	if err != nil {
		log.Error("update status error: ", err)
	}
	return gid, nil
}

func (a *AriaUtil) getStatus(gid string) (string, int) {
	param := []string{"status"}
	res, err := aria2rpc.GetStatus(gid, param)
	if err != nil {
		log.Errorf("get status error. gid:%s, err:%v", gid, err)
	}
	status := fmt.Sprintf("%v", res["status"])
	statusCode := -1
	if temp, ok := statusMap[status]; ok {
		statusCode = temp
	}
	return status, statusCode
}
