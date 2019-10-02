package main


import (
	"flag"
	subcribers "github.com/mikedutuandu/micro_app/micro_teacher/subcribers"
	repo "github.com/mikedutuandu/micro_app/micro_teacher/repositories"
	services "github.com/mikedutuandu/micro_app/micro_teacher/services"
	handlers "github.com/mikedutuandu/micro_app/micro_teacher/handlers"
	publishers "github.com/mikedutuandu/micro_app/micro_teacher/publishers"
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
	services.Init()
	publishers.Init()

	//Init Subscribe for handle event driven data
	subcribers.Init()

	//Init server gRPC for api
	handlers.Init()

	/* End Init  */

}
