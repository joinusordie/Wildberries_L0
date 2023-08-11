package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"

	l0 "github.com/joinusordie/Wildberries_L0"
	"github.com/joinusordie/Wildberries_L0/internal/cache"
	"github.com/joinusordie/Wildberries_L0/internal/handler"
	"github.com/joinusordie/Wildberries_L0/internal/repository"
	"github.com/joinusordie/Wildberries_L0/internal/service"
	subscription "github.com/joinusordie/Wildberries_L0/internal/subscribe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	sc, _ := stan.Connect(viper.GetString("nats.clusterID"), viper.GetString("nats.Subscriber"), stan.NatsURL(viper.GetString("nats.serverURL")))
	defer sc.Close()
	
	repos := repository.NewRepository(db)
	cache := cache.NewCache()
	services := service.NewService(repos, cache)
	handlers := handler.NewHandler(services)

	srv := new(l0.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

	if err = services.AddAllInCache(); err != nil {
		logrus.Fatalf("failed to load cache from db")
	}

	subscription.NewSubscription(sc, services)
	
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
