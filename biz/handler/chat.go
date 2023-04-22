package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"rainbot/biz/model"
)

func SendMsgToChat(ctx context.Context, c *app.RequestContext) {
	req := new(model.MinimaxGPTReq)
	err := c.BindAndValidate(req)
	if err != nil {
		return
	}
	data, err := svcHandler.svc.SendMsg(ctx, *req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, data)
}
