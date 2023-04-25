package larkbot

import (
	"context"
	"fmt"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"rainbot/biz/config"
)

type LarkEvent interface {
}
type larkEventClient struct {
	event *dispatcher.EventDispatcher
}

func NewClient() LarkEvent {
	return &larkEventClient{
		event: dispatcher.NewEventDispatcher(config.LarkBotConfig.Token, ""),
	}
}
func (l *larkEventClient) SendMsg() {

}
func callBackEvent(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	fmt.Println(larkcore.Prettify(event))
	fmt.Println(event.RequestId())
	return nil
}
