package initialize

import (
	"OnlineShop/config"
	"OnlineShop/global"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
)

func ConnectCOS() {
	conf := config.ProjectConfig.COS
	u, _ := url.Parse(conf.Url)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.SecretID,
			SecretKey: conf.SecretKey,
		},
	})
	//client.
	global.GCos = client
	//client.SetDebug(&debug.DebugLogger{})
	if client == nil {
		log.Fatalf("连接COS服务器失败\n")
	}
}
