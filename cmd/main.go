package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	kser "github.com/nartikov/kserver"
	handler "github.com/nartikov/kserver/pkg/handler"
	"github.com/nartikov/kserver/pkg/repository"
	"github.com/nartikov/kserver/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Errorf("error init configs: %s \n", err))
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("pdb.host"),
		Port:     viper.GetString("pdb.port"),
		Username: viper.GetString("pdb.username"),
		DBName:   viper.GetString("pdb.dbname"),
		SSLMode:  viper.GetString("pdb.sslmode"),
		Password: os.Getenv("PDB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(kser.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("run error: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
