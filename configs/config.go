package configs

import "github.com/spf13/viper"

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("port", "8888")
	viper.SetDefault("mongo-db", "")
	viper.SetDefault("mongo_host", "")

}
