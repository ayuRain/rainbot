package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
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
