package main

import (
	"github.com/Heroin-lab/taxi_service.git/internal/app/server"
	"github.com/Heroin-lab/taxi_service.git/pkg/handlers"
	"github.com/Heroin-lab/taxi_service.git/pkg/repositories"
	"github.com/Heroin-lab/taxi_service.git/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initalizing configs: %s", err.Error())
	}

	repos := repositories.NewRepositories()
	services := services.NewServices(repos)
	router := handlers.NewRouter(services)

	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), router.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while starting http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
