package chat

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"io"
	"net/http"
	"rainbot/biz/model"
	"rainbot/biz/util"
)

type MiniMaxGPTService interface {
	SendMsg(ctx context.Context, req model.MinimaxGPTReq) (*model.MinimaxGPTResp, error)
}

type MiniMaxGPTServiceImpl struct {
	GroupID string
	ApiKey  string
	Host    string
}

func (m *MiniMaxGPTServiceImpl) Init(groupID string, apiKey string, host string) {
	m.Host = host
	m.ApiKey = apiKey
	m.GroupID = groupID
}

func (m *MiniMaxGPTServiceImpl) SendMsg(ctx context.Context, req model.MinimaxGPTReq) (*model.MinimaxGPTResp, error) {
	var res = new(model.MinimaxGPTResp)

	dataByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(dataByte)
	reqNew, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s?GroupId=%s", m.Host, m.GroupID), bodyReader)
	if err != nil {
		return nil, err
	}
	reqNew.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.ApiKey))
	reqNew.Header.Add("Content-Type", "application/json")

	resp, err := util.GetDefaultClient().Do(reqNew)
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
