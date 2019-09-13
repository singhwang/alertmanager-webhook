package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/singhwang/alertmanager-webhook/model"
	"github.com/singhwang/alertmanager-webhook/service"
)

func NewWebHookController() webHookController {
	return webHookController{}
}

type webHookController struct {
}

// @Summary PostWebhook
// @Description 接收Alertmanager发送的告警信息.
// @Tags Alertmanager
// @Accept json
// @Produce json
// @Param request body model.Notification true "Alertmanager发送的告警信息."
// @Security ApiKeyAuth
// @Success 200 {object} model.HandleResult ""
// @Failure 400 {object} model.FailureResponse ""
// @Router /api/v1/node/disk/total/usage/ [post]
func (whc *webHookController) Post(c *gin.Context) {
	var notification model.Notification
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailureResponse{
			Error: err.Error(),
		})
		return
	}
	webHookService := service.NewWebHookService()
	handleResult := webHookService.HandleNotification(notification)
	c.JSON(handleResult.Status, handleResult.Result)
}
