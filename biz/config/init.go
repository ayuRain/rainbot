package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var MiniMaxConfig miniMax
var LarkBotConfig larkBot

func InstallBaseConfig() {
	miniMaxBytes, err := os.ReadFile(fmt.Sprintf("./conf/%s.yaml", miniMaxConfigName))
	if err != nil {
		panic(err)
	}
	larkBotBytes, err := os.ReadFile(fmt.Sprintf("./conf/%s.yaml", larkbotConfigName))
	MiniMaxConfig = miniMax{}
	LarkBotConfig = larkBot{}
	err = yaml.Unmarshal(miniMaxBytes, &MiniMaxConfig)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(larkBotBytes, &LarkBotConfig)
	if err != nil {
		panic(err)
	}
}
