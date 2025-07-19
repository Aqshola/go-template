package router

import (
	config_db "go-template/src/config/db"
	config_general "go-template/src/config/general"
	middleware_cors "go-template/src/middleware/cors"
	route_main "go-template/src/router/main"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine, conf config_general.AllConfig, db config_db.Connection) {

	appConfig := conf.AppConfig

	if appConfig.RunMode == "development" {
		route.Use(middleware_cors.AllowCORS())
	}

	mainRoute := route.Group("/")
	route_main.InitMainRoute(mainRoute, db)

}
