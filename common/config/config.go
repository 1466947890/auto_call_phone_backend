package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	MySQL  MySQLConfig  `yaml:"mysql"`
	JWT    JWTConfig    `yaml:"jwt"`
	CORS   CORSConfig   `yaml:"cors"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type MySQLConfig struct {
	DSN      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type JWTConfig struct {
	Secret string `yaml:"secret"`
}

type CORSConfig struct {
	AllowOrigins        []string `yaml:"allow_origins"`
	AllowOriginPrefixes []string `yaml:"allow_origin_prefixes"`
	AllowOriginContains []string `yaml:"allow_origin_contains"`
	AllowMethods        []string `yaml:"allow_methods"`
	AllowHeaders        []string `yaml:"allow_headers"`
	ExposeHeaders       []string `yaml:"expose_headers"`
	AllowCredentials    bool     `yaml:"allow_credentials"`
}

var App Config

func Load(path string) error {
	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	if path == "" {
		path = findDefaultConfigPath()
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config %q: %w", path, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("parse config %q: %w", path, err)
	}
	App = cfg
	return nil
}

func findDefaultConfigPath() string {
	candidates := []string{
		"config.yaml",
		filepath.Join("..", "config.yaml"),
		filepath.Join("/go", "src", "config.yaml"),
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return "config.yaml"
}
