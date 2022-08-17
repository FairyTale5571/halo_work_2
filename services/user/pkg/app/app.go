package app

import (
	"os"

	"github.com/fairytale5571/halo_work_2/services/user/pkg/handler"
	"github.com/fairytale5571/halo_work_2/services/user/pkg/logger"
	"github.com/fairytale5571/halo_work_2/services/user/pkg/server"
)

type App struct {
	Logger *logger.Wrapper
}

func New() (*App, error) {
	log := logger.New("auth_app")

	handlers := handler.NewHandler()
	srv := server.InitServer()
	srv.SetHandler(handlers.InitHandlers())
	go func() {
		if err := srv.Run(); err != nil {
			log.Error(err)
		}
	}()
	log.Infof("auth service is running on port %s", os.Getenv("PORT_AUTH"))

	return &App{
		Logger: log,
	}, nil
}
