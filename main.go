package main

import (
	config_app "go-template/src/config/app"
	config_db "go-template/src/config/db"
	config_general "go-template/src/config/general"
	config_logger "go-template/src/config/logger"
	config_validator "go-template/src/config/validator"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	conf := config_general.InitConfig()
	if conf.AppConfig.RunMode != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	logger := config_logger.NewLogger(conf.AppConfig)
	validator := config_validator.NewValidator()

	server := gin.Default()

	server.Use(gin.Recovery())

	callDb, err := config_db.New(conf)

	if err != nil {
		panic(err)
	}

	serv := config_app.NewServer(server, conf, *callDb, logger, validator)
	config_app.StartService(serv)
}
