package injection

import (
	"github.com/jacobintern/meme_coin/pkg/conf"
	"github.com/jacobintern/meme_coin/repository"
	"github.com/jacobintern/meme_coin/service"
	"go.uber.org/zap"
)

type Injection struct {
	Config          *conf.Config
	ZapLog          *zap.Logger
	MemeCoinService service.IMemeCoinService
}

func New() *Injection {
	// config
	conf := initConfig()

	// clients
	mysql := initMysql(conf)
	logger := initLogger(conf)

	// repository
	memeCoinRepo := repository.NewMemeCoinRepository(mysql)

	// service
	memeCoinService := service.NewMemeCoinService(memeCoinRepo)

	return &Injection{
		Config:          conf,
		ZapLog:          logger,
		MemeCoinService: memeCoinService,
	}
}
