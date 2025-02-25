package injection

import (
	"os"

	"github.com/jacobintern/meme_coin/pkg/conf"
	"github.com/jacobintern/meme_coin/pkg/mysqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initConfig() *conf.Config {
	config, err := conf.NewClient()
	if err != nil {
		panic(err)
	}

	return config
}

func initMysql(conf *conf.Config) *mysqlx.Client {
	return mysqlx.NewClient(conf)
}

func initLogger(conf *conf.Config) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	writeSyncer := zapcore.AddSync(os.Stdout)

	var level zapcore.Level
	switch conf.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warning":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core)
	if logger == nil {
		panic("zap log init failed.")
	}

	return logger
}
