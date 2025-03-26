package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env                  string               `yaml:"env"`
	EmailValidatorConfig EmailValidatorConfig `yaml:"emailvalidator"`
}

type EmailValidatorConfig struct {
	Port int `yaml:"port"`
}

func MustLoad() *Config {
	var cfg Config

	env := os.Getenv("env")

	switch env {
	case "dev":
		err := cleanenv.ReadConfig("./config/dev.yaml", &cfg)
		if err != nil {
			panic("could't read config: " + err.Error())
		}

	case "prod":
		err := cleanenv.ReadConfig("./config/prod.yaml", &cfg)
		if err != nil {
			panic("could't read config: " + err.Error())
		}

	default:
		err := cleanenv.ReadConfig("./emailvalidator/config/local.yaml", &cfg)
		if err != nil {
			panic("could't read config: " + err.Error())
		}
	}

	return &cfg
}
