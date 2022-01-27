package main

import (
	"fmt"
	"go-trading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
