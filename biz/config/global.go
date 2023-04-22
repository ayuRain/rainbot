package config

const defaultConfigName = "global"

type global struct {
	GroupID string `json:"group_id" yaml:"group_id"`
	Secret  string `json:"secret" yaml:"secret"`
	Host    string `json:"host" yaml:"host"`
}
