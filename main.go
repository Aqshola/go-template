package main

import (
	config_app "go-template/src/config/app"
	config_db "go-template/src/config/db"
	config_general "go-template/src/config/general"

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

	server := gin.Default()

	server.Use(gin.Recovery())

	callDb, err := config_db.New(conf)

	if err != nil {
		panic(err)
	}

	serv := config_app.NewServer(server, conf, *callDb)
	config_app.StartService(serv)
}
