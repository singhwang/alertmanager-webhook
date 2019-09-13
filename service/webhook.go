package service

import (
	"net/http"

	"github.com/golang/glog"

	"github.com/singhwang/alertmanager-webhook/model"
)

func NewWebHookService() webHookService {
	return webHookService{}
}

type webHookService struct {
}

func (s *webHookService) HandleNotification(notification model.Notification) model.HandleResult {
	glog.Info("notification.Version : ", notification.Version)
	glog.Info("notification.GroupKey : ", notification.GroupKey)
	glog.Info("notification.Status : ", notification.Status)
	glog.Info("notification.Receiver : ", notification.Receiver)
	glog.Info("notification.GroupLabels : ", notification.GroupLabels)
	glog.Info("notification.CommonLabels : ", notification.CommonLabels)
	glog.Info("notification.CommonAnnotations : ", notification.CommonAnnotations)
	glog.Info("notification.ExternalURL : ", notification.ExternalURL)
	alerts := notification.Alerts
	for i := 0; i < len(alerts); i++ {
		alert := alerts[i]
		glog.Info("alert.Labels : ", alert.Labels)
		glog.Info("alert.Annotations : ", alert.Annotations)
		glog.Info("alert.StartsAt : ", alert.StartsAt)
		glog.Info("alert.EndsAt : ", alert.EndsAt)
	}

	// TODO 处理告警信息（例如：发邮件、存入数据库等等。）

	return model.HandleResult{
		Status: http.StatusOK,
		Result: "Received alarm information successfully.",
	}
}
