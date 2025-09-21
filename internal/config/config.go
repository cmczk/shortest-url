package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"prod"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:3000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	fmt.Println("start config reading")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("cannot find CONFIG_PATH environment variable")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot parse config file: %s", err.Error())
	}

	return &cfg
}
