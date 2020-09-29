package main

import (
	"fmt"
	"gotrading/app/controllers"
	"gotrading/app/models"
	"gotrading/config"
	"gotrading/utils"
	"log"
	"time"
)

func main() {
	df, _ := models.GetAllCandle(config.Config.ProductCode, time.Minute,365)
	fmt.Printf("%+v\n", df.OptimizeParams())
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
