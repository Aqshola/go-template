package controller_main

import (
	common_parsing "go-template/src/common/parsing"

	config_db "go-template/src/config/db"
	repository_main "go-template/src/repository/main/testTable"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type IMainController interface {
	GetMain(ctx *gin.Context)
	GetDetailMain(ctx *gin.Context)
}

type MainController struct {
	repoMain  repository_main.ITestTableRepository
	logger    *logrus.Logger
	validator *validator.Validate
}

func (c *MainController) GetMain(ctx *gin.Context) {
	listData, err := c.repoMain.GetListTestTable()
	if err != nil {
		common_parsing.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	responseData := TransformGetMainResponse(listData)
	common_parsing.JSONResponse(ctx, http.StatusOK, "HALO", responseData)
}

func (c *MainController) GetDetailMain(ctx *gin.Context) {
	var requestParam = GetDetailRequestParam{}
	err := ctx.ShouldBindQuery(&requestParam)
	if err != nil {
		c.logger.Error("GetDetailMain " + err.Error())
		common_parsing.JSONResponse(ctx, http.StatusBadRequest, "Invalid Query Parameters ", nil)
		return
	}

	err = c.validator.Struct(requestParam)
	if err != nil {
		common_parsing.JSONResponse(ctx, http.StatusBadRequest, "Invalid Query Parameters "+err.Error(), nil)
		return
	}

	dataDetail, err := c.repoMain.GetDetailTestTable(requestParam.Id)
	if err != nil {
		common_parsing.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	responseData := TransformGetDetailMainResponse(*dataDetail)
	common_parsing.JSONResponse(ctx, http.StatusBadRequest, "SUCCESS", responseData)
}

func NewMainController(db config_db.Connection, logger *logrus.Logger, validator *validator.Validate) IMainController {
	repoMain := repository_main.NewMainRepository(db.MySQL)

	return &MainController{
		repoMain:  repoMain,
		logger:    logger,
		validator: validator,
	}
}
