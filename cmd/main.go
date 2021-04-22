package main

import (
	"log"

	kser "github.com/nartikov/kserver"
	handler "github.com/nartikov/kserver/pkg/handler"
	"github.com/nartikov/kserver/pkg/repository"
	"github.com/nartikov/kserver/pkg/service"
)

func main()  {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv:=new(kser.Server)
	if err:=srv.Run("8080",handlers.InitRoutes()); err!=nil{
		log.Fatalf("run error: %s", err.Error())
	}

}