package conf

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jacobintern/meme_coin/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `validate:"required"`
	}
	Log struct {
		Level string
	}
	Http struct {
		TimeoutSec int
	}
	Database struct {
		MySql struct {
			Database string `validate:"required"`
			User     string `validate:"required"`
			Password string `validate:"required"`
			Host     string `validate:"required"`
			Port     string `validate:"required"`
		}
	}
	Cors struct {
		ALLOW struct {
			ALLOrigins  bool
			Credentials bool
			Methods     []string `validate:"required"`
			Headers     []string `validate:"required"`
		}
		MaxAge int64 `validate:"required"`
	}
}

func NewClient() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, errors.ReadConfigFatal.Wrap(err)
	}

	v.AutomaticEnv()

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "unmarshal value to config struct failed.")
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, errors.Wrap(err, "config validation failed.")
	}

	return &config, nil
}
