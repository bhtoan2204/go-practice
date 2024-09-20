package initialize

import (
	"fmt"

	"github.com/bhtoan2204/go-practice.git/global"
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
