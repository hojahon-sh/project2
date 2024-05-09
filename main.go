package main

import (
	"fmt"

	"github.com/project2/repository"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello!")

	conf := viper.New()
	conf.SetConfigFile(".env")

	if err := conf.ReadInConfig(); err != nil {
		fmt.Println("Error reading env file", err)
		return

	}

	err := repository.Connect(conf)
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
}
