package main

import (
	"flag"
	"tiramisu/config"
	"tiramisu/model"

	"github.com/NoahShen/aria2rpc"
)

var (
	taskModel model.TaskStruct
	ariaUtil  AriaUtil
)

func init() {
	flag.Parse()
	flag.Lookup("alsologtostderr").Value.Set("true")
	aria2rpc.RpcUrl = config.RPC_URL
}

func main() {

}
