package main

import (
	"fmt"
	"os"

	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	sc, _ := stan.Connect(viper.GetString("nats.clusterID"), viper.GetString("nats.producer"), stan.NatsURL(viper.GetString("nats.serverURL")))
	defer sc.Close()
	jsonFile, err := os.ReadFile("model.json")
	if err != nil {
		logrus.Fatalf("error read json file: %s", err.Error())
	}

	sc.Publish(viper.GetString("nats.subject"), jsonFile)
	var a string
	fmt.Scan(a)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}