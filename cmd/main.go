package main

import (
	"github.com/MatheusNP/ropasclispo/internal/interface/rest"
	"github.com/MatheusNP/ropasclispo/internal/service"
)

func main() {

	playService := service.NewPlayService()

	controller := rest.Controllers{
		PlayController: rest.NewPlayController(playService),
	}

	controller.NewServer()
}
