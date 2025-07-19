package controller_main

import (
	common_parsing "go-template/src/common/parsing"
	config_db "go-template/src/config/db"
	repository_main "go-template/src/repository/main/testTable"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IMainController interface {
	GetMain(ctx *gin.Context)
	GetDetailMain(ctx *gin.Context)
}

type MainController struct {
	repoMain repository_main.ITestTableRepository
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
	idMain := ctx.Query("id")
	parsedIdMain, err := strconv.Atoi(idMain)
	if err != nil {
		common_parsing.JSONResponse(ctx, http.StatusBadRequest, "Invalid project ID", nil)
		return
	}

	dataDetail, err := c.repoMain.GetDetailTestTable(int(parsedIdMain))
	if err != nil {
		common_parsing.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	responseData := TransformGetDetailMainResponse(*dataDetail)
	common_parsing.JSONResponse(ctx, http.StatusBadRequest, "SUCCESS", responseData)
}

func NewMainController(db config_db.Connection) IMainController {
	repoMain := repository_main.NewMainRepository(db.MySQL)
	return &MainController{
		repoMain: repoMain,
	}
}
