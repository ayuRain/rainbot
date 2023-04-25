package larkbot

import (
	lark_hertz "github.com/hertz-contrib/lark-hertz"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
)

type LarkBotService interface {
}

type LarkBotServiceImpl struct {
	Cli *lark.Client
}

func (l *LarkBotServiceImpl) Init(appID string, secret string) {
	l.Cli = lark.NewClient(appID, secret)
}

func (l *LarkBotServiceImpl) loop() {
	handler := dispatcher.NewEventDispatcher("", "").OnP2MessageReceiveV1()
	lark_hertz.NewEventHandlerFunc(handler)
}

func (l *LarkBotServiceImpl) Revice() {

}
