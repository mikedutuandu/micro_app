package main


import (
	"flag"
	subcribers "github.com/mikedutuandu/micro_app/micro_notification/subcribers"
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

	//Init Subscribe for handle event driven data
	subcribers.Init()

	/* End Init  */

}
