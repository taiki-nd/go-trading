package main

import (
	"fmt"
	"go-trading/bitflyer"
	"go-trading/config"
	"go-trading/utils"
	"time"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker, _ := apiClient.GetTicker("BTC_USD")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}
