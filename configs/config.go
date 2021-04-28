package configs

import "github.com/spf13/viper"

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("port", "8888")
	viper.SetDefault("mongo_database", "carlos")
	viper.SetDefault("mongo_host", "mongodb://localhost:27017")

}
