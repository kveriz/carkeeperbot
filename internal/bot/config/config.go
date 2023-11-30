package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Token       string `yaml:"token"`
	Lang        string `yaml:"lang"`
	SQL         struct {
		DBType     string `yaml:"driver"`
		DBUser     string `yaml:"user"`
		DBPassword string `yaml:"password"`
		DBHost     string `yaml:"host"`
		DB         string `yaml:"database"`
		Mode       string `yaml:"mode"`
		SqlitePath string `yaml:"sqlitePath"`
	}
}

func New(file string) *Config {
	config := &Config{}

	confFile, err := os.Open(file)
	if err != nil {
		log.Printf("Open config file error: %v", err)
	}
	defer confFile.Close()

	dec := yaml.NewDecoder(confFile)
	if err := dec.Decode(&config); err != nil {
		return nil
	}

	return config
}
