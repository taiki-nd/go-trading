package main

import (
	"fmt"
	"go-trading/app/models"
	"go-trading/config"
	"go-trading/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	/*
		apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

		tickerChannel := make(chan bitflyer.Ticker)
		go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
		for ticker := range tickerChannel {
			fmt.Println(ticker)
			fmt.Println(ticker.GetMidPrice())
			fmt.Println(ticker.DateTime())
			fmt.Println(ticker.TruncateDateTime(time.Second))
			fmt.Println(ticker.TruncateDateTime(time.Minute))
			fmt.Println(ticker.TruncateDateTime(time.Hour))
		}
	*/
	fmt.Println(models.DbConnection)
}
