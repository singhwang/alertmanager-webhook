package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/singhwang/alertmanager-webhook/controller"
	_ "github.com/singhwang/alertmanager-webhook/docs"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func NewWebHookRouter() webHookRouter {
	return webHookRouter{}
}

type webHookRouter struct {
}

func (r *webHookRouter) Run() error {
	domain := viper.GetString("domain")
	port := viper.GetString("port")

	router := gin.Default()
	url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", domain, port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	webHookController := controller.NewWebHookController()
	router.POST("/webhook", webHookController.Post)
	return router.Run()
}
