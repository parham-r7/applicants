package main

import (
	"applicants/services"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main()  {
	services.Run()
}
