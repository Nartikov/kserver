package main

import (
	"log"

	kser "github.com/nartikov/kserver"
	handler "github.com/nartikov/kserver/pkg/handler"
	"github.com/nartikov/kserver/pkg/repository"
	"github.com/nartikov/kserver/pkg/service"
	"github.com/spf13/viper"
)

func main()  {
	if err:=initConfig(); err !=nil{
		log.Fatalf("error init configs: %s", err.Error())
	}



	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv:=new(kser.Server)
	if err:=srv.Run(viper.GetString("port"),handlers.InitRoutes()); err!=nil{
		log.Fatalf("run error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}