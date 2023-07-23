package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvFilePath = ".env"
)

type Config struct {
	HTTP struct {
		IP   string `env:"HTTP_IP"`
		Port uint16 `env:"HTTP_PORT"`
	}

	GRPC struct {
		IP                   string `env:"GRPC_IP"`
		Port                 uint16 `env:"GRPC_PORT"`
		MaxConcurrentStreams uint32 `env:"GRPC_MAX_CONCURRENT_STREAMS"`
	}

	Redis struct {
		IP       string `env:"REDIS_IP"`
		Port     uint16 `env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD"`
		DB       int    `env:"REDIS_DB"`
	}

	Cors struct {
		AllowOrigin  string `env:"CORS_ALLOW_ORIGIN"`
		AllowMethods string `env:"CORS_ALLOW_METHODS"`
	}
}

func ReadEnvFile() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(EnvFilePath, &cfg)

	return &cfg, err
}
