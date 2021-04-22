package main

import (
	"log"
	handler "github.com/nartikov/kserver/pkg/handler"
	kser "github.com/nartikov/kserver"
)

func main()  {
	handlers:=new(handler.Handler)
	srv:=new(kser.Server)
	if err:=srv.Run("8000",handlers.InitRoutes()); err!=nil{
		log.Fatalf("run error: %s", err.Error())
	}

}