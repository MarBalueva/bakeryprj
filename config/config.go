package config

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Config — структура для конфигурации из YAML
type Config struct {
	App struct {
		Env  string `yaml:"env"`
		Port int    `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

var Cfg Config

// LoadEnv загружает .env файл (если есть)
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

// GetEnv возвращает значение переменной окружения или дефолт
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// LoadYAMLConfig загружает конфиг из YAML файла
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

// LoadConfig инициализация конфигурации из .env и yaml
func LoadConfig() {
	LoadEnv()

	// Загружаем yaml, игнорируем ошибку
	if err := LoadYAMLConfig("config.yaml"); err != nil {
		log.Println("Cannot load config.yaml:", err)
	}

	// Перезаписываем поля из переменных окружения, если есть
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
