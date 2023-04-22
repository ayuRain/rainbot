package service

import (
	"rainbot/biz/config"
	"rainbot/biz/service/chat"
)

type Service interface {
	chat.MiniMaxGPTService
}

type serviceImpl struct {
	chat.MiniMaxGPTServiceImpl
}

func (s *serviceImpl) Init() {
	s.MiniMaxGPTServiceImpl.Init(config.BaseConfig.GroupID, config.BaseConfig.Secret, config.BaseConfig.Host)
}

func Build() Service {
	svc := &serviceImpl{}
	svc.Init()
	return svc
}
