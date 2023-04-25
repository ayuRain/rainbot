package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	lark_hertz "github.com/hertz-contrib/lark-hertz"
	"net/http"
	"rainbot/biz/model"
)

func Verify(ctx context.Context, c *app.RequestContext) {
	req := new(model.VerifyReq)
	err := c.BindAndValidate(req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &model.VerifyResp{Challenge: req.Challenge})
}

func ReceiveEvent(ctx context.Context, c *app.RequestContext) {
	lark_hertz.NewEventHandlerFunc("")
}
