package main

import (
	"github.com/NoahShen/aria2rpc"
	log "github.com/golang/glog"
	"tiarmisu/config"
)

type AriaUtil struct{}

func init() {
	aria2rpc.RpcUrl = config.RPC_URL
}

func (a *AriaUtil) addUri(url string) {
	param := make(map[string]interface{})
	res, err := aria2rpc.AddUri(url, param)
	if err != nil {
		log.Error("add uri error ", err)
	}
	log.Info("res", res)
}
