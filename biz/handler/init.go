package handler

import "rainbot/biz/service"

var svcHandler *handler

type handler struct {
	svc service.Service
}

func BindService() {
	svcHandler = &handler{
		svc: service.Build(),
	}
}
