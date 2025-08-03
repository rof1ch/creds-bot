package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ENV      string `env:"ENV"`
	BotToken string `env:"BOT_TOKEN"`
	DBPath   string `env:"DB_PATH"`
}

func LoadConfig() (*Config, error) {
	envPath, configPath := fetchPaths()

	if envPath == "" {
		return nil, fmt.Errorf("'.env' file path is empty")
	}

	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("no .env file found")
	}

	cfg, err := LoadPath(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadPath(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("can not read config: %w", err)
	}

	return &cfg, nil
}

func fetchPaths() (string, string) {
	var envPath, configPath string

	flag.StringVar(&envPath, "env", "", "path to '.env' file")
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if envPath == "" {
		envPath = os.Getenv("ENV_PATH")
	}

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	return envPath, configPath
}
