package configs

import _ "embed"

//go:embed config.yaml

var ConfigYaml []byte

func GetConfigYaml() []byte {
	return ConfigYaml
}
