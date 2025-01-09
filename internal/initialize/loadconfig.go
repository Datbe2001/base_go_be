package initialize

import (
	"base_go_be/global"
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	vp := viper.New()
	vp.AddConfigPath("./config")
	vp.SetConfigName("local")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//fmt.Println("Server Port::", vp.GetString("security.jwt.key"))

	//var config Config
	if err := vp.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unmarshal err:", err)
	}
}
