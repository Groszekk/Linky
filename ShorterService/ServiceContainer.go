package main

import (
	"Linky/ShorterService/controllers"
	"Linky/ShorterService/database"
	"Linky/ShorterService/infrastructures"
	"Linky/ShorterService/repositories"
	"Linky/ShorterService/services"
	"sync"
)

type IServiceContainer interface {
	InjectController() controllers.LinkController
}

type kernel struct{}

var (
	k             *kernel
	containerOnce sync.Once
)

func (k *kernel) InjectController() controllers.LinkController {
	db := database.InitConnection()
	sqlHandler := &infrastructures.SQLHandler{}
	sqlHandler.Conn = db

	linkRepository := &repositories.LinkRepository{sqlHandler}
	linkService := &services.LinkService{&repositories.LinkRepositoryWithCircuitBreaker{linkRepository}}
	linkController := controllers.LinkController{linkService}

	return linkController
}

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}

	return k
}
