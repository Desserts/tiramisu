package main

import (
	"tiramisu/model"
	log "github.com/golang/glog"
)

var(
	taskModel model.TaskStruct
	ariaUtil AriaUtil
)

func getOneToDownload(){
	task,err := taskModel.GetOneWait()
	if err != nil{
		log.Error("get task error: ",err)
	}
	switch task.Types {
	case model.TASK_TYPE_NORMAL:
		ariaUtil.addUri(task.Url)
	}


}

func main(){

}
