package config

import (
	"bakeryapp/logger"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Env  string `yaml:"env"`
		Port int    `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
	} `yaml:"database"`
}

var Cfg Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Warn("No .env file found, using system environment variables")
	}
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadYAMLConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig() {
	LoadEnv()

	if err := LoadYAMLConfig("config.yaml"); err != nil {
		logger.Log.Error("Cannot load config.yaml:", err)
	}

	Cfg.App.Env = GetEnv("APP_ENV", Cfg.App.Env)
	if portStr := GetEnv("SERVER_PORT", ""); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			Cfg.App.Port = port
		}
	}
	Cfg.Database.Host = GetEnv("DB_HOST", Cfg.Database.Host)
	Cfg.Database.User = GetEnv("DB_USER", Cfg.Database.User)
	Cfg.Database.Password = GetEnv("DB_PASSWORD", Cfg.Database.Password)
}
