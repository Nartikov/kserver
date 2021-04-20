package main

import (
	"log"

	kser "github.com/nartikov/kserver"
)

func main()  {
	srv:=new(kser.Server)
	if err:=srv.Run("8000"); err!=nil{
		log.Fatalf("run error: %s", err.Error())
	}

}