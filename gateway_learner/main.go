package main

import (
	"flag"
	"github.com/mikedutuandu/micro_app/gateway_learner/handlers"
	"github.com/mikedutuandu/micro_app/gateway_learner/services"
	"github.com/spf13/viper"
)

func initConfig() {
	var configFile = flag.String("config", "./configs/config-dev.yml", "Path to config file.")
	flag.Parse()
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic("Error loading config.json")
	}
}

func main() {
	initConfig()
	//Init service
	services.Init()

	//Init http server
	handlers.Init()
}
