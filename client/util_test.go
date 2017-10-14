package main

import (
	log "github.com/golang/glog"
	"testing"
	"time"
)

func TestAddUrlNormal(t *testing.T) {
	url1 := "https://dldir1.qq.com/qqfile/QQforMac/QQ_V6.1.0.dmg"
	//url2 := "https://www.baidu.com"
	id := 1
	gid, _ := ariaUtil.addUriNormal(id, url1)
	for {
		time.Sleep(3 * time.Second)
		res, code := ariaUtil.getStatus(gid)
		log.Info(res)
		taskModel.UpdateGidAndStatus(id, gid, code)
	}
}

func TestDownloadNormal(t *testing.T) {
	data := []struct {
		url string
	}{
		{"magnet:?xt=urn:btih:A228E6AEF537F1C7CE246A6DF8DBDE94C9581EC8"},
	}
	for _, v := range data {
		ariaUtil.addUriNormal(1, v.url)
	}
}

func TestGetStatus(t *testing.T) {
	url := "https://nginx.org/download/nginx-1.12.1.tar.gz"
	ariaUtil.addUriNormal(1, url)
}
