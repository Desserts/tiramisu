package main

import "testing"

func TestInit(t *testing.T) {
	ariaUtil.addUri("https://nginx.org/download/nginx-1.12.1.tar.gz")
	//ariaUtil.AddUri("https://www.baidu.com")
}

func TestDownloadNormal(t *testing.T) {
	data := []struct {
		url string
	}{
		{"magnet:?xt=urn:btih:A228E6AEF537F1C7CE246A6DF8DBDE94C9581EC8"},
	}
	for _, v := range data {
		ariaUtil.addUri(v.url)
	}
}
