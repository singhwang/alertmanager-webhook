package main

import (
	"flag"
	"os"

	"github.com/golang/glog"

	"github.com/singhwang/alertmanager-webhook/router"
	"github.com/singhwang/alertmanager-webhook/util/config"
)

const (
	GlogKey    = "logtostderr"
	GlogValue  = "true"
	ConfigPath = "config/alertmanager-webhook.yaml"
	ErrorCode  = 1
)


// @title alertmanager-webhook服务
// @version 1.0
// @description 用于接收alertmanager告警信息的webhook服务.

// @contact.name 王欣
// @contact.email wangxin_0611@126.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

func main() {
	flag.Parse()
	err := flag.Set(GlogKey, GlogValue)
	if err != nil {
		glog.Error("Init glog occur a error. It is ", err, " .")
		os.Exit(ErrorCode)
	}

	awc := config.NewConfig()
	err = awc.Init(ConfigPath)
	if err != nil {
		glog.Error("Init config occur a error. It is ", err, " .")
		os.Exit(ErrorCode)
	}

	webHookRouter := router.NewWebHookRouter()
	err = webHookRouter.Run()
	if err != nil {
		glog.Error("Run webHookRouter occur a error. It is ", err, ".")
		os.Exit(ErrorCode)
	}
}
