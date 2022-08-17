package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/fairytale5571/halo_work_2/services/proxy/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal("error create app: ", err)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	a.Logger.Info("shutdown application")
}
