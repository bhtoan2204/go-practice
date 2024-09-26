package initialize

import (
	"fmt"

	"gateway-service/global"

	"github.com/spf13/viper"
)

func InitViper() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode configuration: %s", err))
	}
}
