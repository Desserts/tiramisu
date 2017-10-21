package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"tiramisu/model"

	log "github.com/golang/glog"
)

type ResDataStruct struct {
	Succ bool
	Msg  string
}

var (
	taskModel *model.TaskStruct
)

func index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":   "title",
		"Content": "aaaa",
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, data)
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//r.ParseForm()
	url := r.PostFormValue("url")
	log.Info("get url: ", url)
	err := taskModel.AddTask(url, model.TASK_TYPE_NORMAL)

	res := ResDataStruct{
		Succ: true,
	}

	if err != nil {
		log.Error("add task error: ", err)
		res.Msg = err.Error()
		resStr, err := json.Marshal(res)
		if err != nil {
			resStr = []byte(err.Error())
		}
		w.Write(resStr)
	}
}
