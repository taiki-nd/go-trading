package main

import (
	"fmt"
	"go-trading/config"
	"go-trading/utils"
	"log"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	log.Println("test")
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
