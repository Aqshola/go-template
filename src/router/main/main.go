package route_main

import (
	config_db "go-template/src/config/db"
	controller_main "go-template/src/controller/main"

	"github.com/gin-gonic/gin"
)

func InitMainRoute(route *gin.RouterGroup, db config_db.Connection) {
	mainController := controller_main.NewMainController(db)

	route.GET("/", mainController.GetMain)
	route.GET("/detail", mainController.GetDetailMain)
}
