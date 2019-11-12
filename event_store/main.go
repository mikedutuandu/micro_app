package main


import (
	"flag"
	publishers "github.com/mikedutuandu/micro_app/event_store/publishers"
	repo "github.com/mikedutuandu/micro_app/event_store/repositories"
	handlers "github.com/mikedutuandu/micro_app/event_store/handlers"
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

	/* Init  */
	initConfig()
	repo.Init()

	publishers.Init()

	//Init server gRPC for api
	handlers.Init()

	/* End Init  */

}
