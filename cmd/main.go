package main

import (
	"log"
	"github.com/nartikov/kserver/pkg/hendler"
	kser "github.com/nartikov/kserver"
)

func main()  {
	hendlers:=new(hendler.Hendler)
	srv:=new(kser.Server)
	if err:=srv.Run("8000",hendlers.InitRoutes()); err!=nil{
		log.Fatalf("run error: %s", err.Error())
	}

}