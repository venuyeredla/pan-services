package config

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Http struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		ReadTimeout     int    `yaml:"read_timeout"`
		WriteTimeout    int    `yaml:"write_timeout"`
		IdleTimeout     int    `yaml:"idle_timeout"`
		ShutdownTimeout int    `yaml:"shutdown_timeout"`
	} `yaml:"http"`
}

func NewConfig(configData []byte) (*Config, error) {

	var config Config
	err := yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
		return nil, err
	}
	log.Printf("Config loaded successfully: %+v", config)
	return &config, nil
}
