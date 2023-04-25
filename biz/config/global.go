package config

const miniMaxConfigName = "minimax"
const larkbotConfigName = "larkbot"

type miniMax struct {
	GroupID string `json:"group_id" yaml:"group_id"`
	Secret  string `json:"secret" yaml:"secret"`
	Host    string `json:"host" yaml:"host"`
}

type larkBot struct {
	AppID  string `json:"app_id" yaml:"app_id"`
	Secret string `json:"secret" yaml:"secret"`
	Token  string `json:"token" yaml:"token"`
}
