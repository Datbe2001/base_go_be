package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `map_structure:"port"`
	} `map_structure:"server"`
	Databases []struct {
		User     string `map_structure:"user"`
		Password string `map_structure:"password"`
		Host     string `map_structure:"host"`
	} `map_structure:"databases"`
	Security struct {
		Jwt struct {
			SecretKey string `map_structure:"secretKey"`
		} `map_structure:"jwt"`
	}
}

func main() {
	vp := viper.New()
	vp.AddConfigPath("./config")
	vp.SetConfigName("local")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//fmt.Println("Server Port::", vp.GetString("security.jwt.key"))

	var config Config
	if err := vp.Unmarshal(&config); err != nil {
		fmt.Println("Unmarshal err:", err)
	}

	fmt.Println("Config Port:", config.Server.Port)
	fmt.Println("JWT:", config.Security.Jwt.SecretKey)
	for _, database := range config.Databases {
		fmt.Printf("User: %s, password: %s, host: %s", database.User, database.Password, database.Host)
	}

}
