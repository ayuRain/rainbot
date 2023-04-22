package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var BaseConfig global

func InstallBaseConfig() {
	dataBytes, err := os.ReadFile(fmt.Sprintf("./conf/%s.yaml", defaultConfigName))
	if err != nil {
		panic(err)
	}
	BaseConfig = global{}
	err = yaml.Unmarshal(dataBytes, &BaseConfig)
	if err != nil {
		panic(err)
	}
}
