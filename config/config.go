package config

import (
	"gitlab.com/bongerka/vault"
	"gitlab.com/piorun102/lg"
)

type (
	Config struct {
		User    UserConfig    `json:"users"`
		Service ServiceConfig `json:"services"`
		Shared  SharedConfig  `json:"shared"`
	}
	UserConfig struct {
		Postgres PostgresUser `json:"postgres"`
	}

	PostgresUser struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}

	ServiceConfig struct {
		Server     ServerConfig     `json:"server"`
		Processing ProcessingConfig `json:"processing"`
	}

	ProcessingConfig struct {
		ApiPublic  string
		ApiPrivate string
		Url        string
	}

	ServerConfig struct {
		Host string
		Port string
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Db       string
		Password string
		Ssl      string
		User     string
	}

	SharedConfig struct {
		Nats     NatsShared     `json:"nats"`
		Postgres PostgresShared `json:"postgres"`
		Logger   LoggerShared   `json:"logger"`
		Jaeger   JaegerShared   `json:"jaeger"`
	}
	JaegerShared struct {
		Url string
	}
	NatsShared struct {
		Url      string
		User     string
		Password string
	}

	PostgresShared struct {
		Host string `json:"url"`
		Port string `json:"port"`
		Db   string `json:"db"`
		Ssl  string `json:"ssl"`
	}

	LoggerShared struct {
		Address        string `json:"url"`
		MetricsAddress string `json:"metrics-url"`
	}
)

func Load() *Config {
	lg.Trace("Loading config from vault started")
	var cfg Config
	vault.GetConfig(&cfg)
	lg.Trace("Loading config from vault completed")
	return &cfg
}
