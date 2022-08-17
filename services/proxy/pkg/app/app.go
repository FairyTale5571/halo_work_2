package app

import (
	"github.com/fairytale5571/halo_work_2/services/proxy/pkg/logger"
	"github.com/fairytale5571/halo_work_2/services/proxy/pkg/server"
)

type App struct {
	Logger *logger.Wrapper
}

func New() (*App, error) {
	log := logger.New("proxy_app")

	rout := server.New()
	go func() {
		err := rout.Run()
		if err != nil {
			log.Fatalf("error run server: %v", err)
		}
	}()

	return &App{
		Logger: log,
	}, nil
}
