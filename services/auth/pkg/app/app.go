package app

import (
	"os"

	"github.com/fairytale5571/halo_work_2/services/auth/pkg/handler"
	"github.com/fairytale5571/halo_work_2/services/auth/pkg/logger"
	"github.com/fairytale5571/halo_work_2/services/auth/pkg/server"
	"github.com/fairytale5571/halo_work_2/services/auth/pkg/service"
)

type App struct {
	Logger *logger.Wrapper
}

func New() (*App, error) {
	log := logger.New("auth_app")

	services := service.NewService()
	handlers := handler.NewHandler(services)
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
