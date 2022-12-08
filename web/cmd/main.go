package main

import (
	"coursework/web/internal/broker"
	handler2 "coursework/web/internal/handler"
	"coursework/web/internal/repositories"
	"coursework/web/internal/services"
	"github.com/siruspen/logrus"
	"github.com/spf13/viper"
	"log"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	rabbit := broker.NewBroker(
		viper.GetString("rabbit_user"),
		viper.GetString("rabbit_password"),
		viper.GetString("rabbit_pot"))
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := repositories.NewPostgresDb(viper.GetString("db_url"))
	if err != nil {
		logrus.Fatal(err)
	}
	repository := repositories.NewRepository(db)
	service := services.NewService(repository, rabbit)
	handler := handler2.NewHandler(service)
	srv := new(Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoues()); err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("web/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
