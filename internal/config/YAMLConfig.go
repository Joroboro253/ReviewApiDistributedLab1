package config

type YAMLConfig struct {
	DB struct {
		URL string `yaml:"url"`
	} `yaml:"db"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}
