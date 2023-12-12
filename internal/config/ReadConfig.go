package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadConfig(path string) (YAMLConfig, error) {
	var cfg YAMLConfig

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
