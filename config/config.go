package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		Log      `yaml:"logger"`
		MYSQL    `yaml:"mysql"`
		AWS      `yaml:"aws"`
		RcClient `yaml:"rc_client"`
		Partner  `yaml:"partner"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	MYSQL struct {
		OpenConnMax int `env-required:"true" yaml:"open_conn_max" env:"MYSQL_OPEN_CONN_MAX"`
		IdleConnMax int `env-required:"true" yaml:"idle_conn_max" env:"MYSQL_IDLE_CONN_MAX"`
		LifeConnMax int `env-required:"true" yaml:"life_conn_max" env:"MYSQL_LIFE_CONN_MAX"`
	}

	AWS struct {
		RegionName string `env-required:"true" yaml:"region_name"    env:"AWS_REGION_NAME"`
		SecretName string `env-required:"true" yaml:"secret_name" env:"SECRETS_MANAGER"`
	}

	RcClient struct {
		Timeout int `env-required:"true" yaml:"timeout"    env:"RC_CLIENT_TIMEOUT"`
	}

	Partner struct {
		ApiKey string `env-required:"true" yaml:"apiKey"    env:"PARTNER_API_KEY"`
		URL    string `env-required:"true" yaml:"url"    env:"PARTNER_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
