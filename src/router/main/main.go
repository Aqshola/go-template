package route_main

import (
	config_db "go-template/src/config/db"
	controller_main "go-template/src/controller/main"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func InitMainRoute(route *gin.RouterGroup, db config_db.Connection, logger *logrus.Logger, validator *validator.Validate) {
	mainController := controller_main.NewMainController(db, logger, validator)

	route.GET("/", mainController.GetMain)
	route.GET("/detail", mainController.GetDetailMain)
}
