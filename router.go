// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "rainbot/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.POST("/send", handler.SendMsgToChat)
	r.POST("", handler.Verify)
	// your code ...
}