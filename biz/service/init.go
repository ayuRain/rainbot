package service

import (
	"rainbot/biz/config"
	"rainbot/biz/service/chat"
	"rainbot/biz/service/larkbot"
)

type Service interface {
	chat.MiniMaxGPTService
	larkbot.LarkBotService
}

type serviceImpl struct {
	chat.MiniMaxGPTServiceImpl
	larkbot.LarkBotServiceImpl
}

func (s *serviceImpl) Init() {
	s.MiniMaxGPTServiceImpl.Init(config.MiniMaxConfig.GroupID, config.MiniMaxConfig.Secret, config.MiniMaxConfig.Host)
	s.LarkBotServiceImpl.Init(config.LarkBotConfig.AppID, config.LarkBotConfig.Secret)
}

func Build() Service {
	svc := &serviceImpl{}
	svc.Init()
	return svc
}
