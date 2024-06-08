package handlers

import (
	"crud/common/logger"
	"crud/pkg/entity/requestDto"
	"crud/pkg/entity/responseDto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func validateGetEmployee(ctx *gin.Context, request *requestDto.GetEmployeeRequest) error {
	employeeId := ctx.Param("employeeId")
	var err error
	request.EmployeeId, err = strconv.Atoi(employeeId)
	return err
}

func (app *ApiHandler) GetEmpoyee(ctx *gin.Context) {
	var request requestDto.GetEmployeeRequest
	err := validateGetEmployee(ctx, &request)
	if err != nil {
		logger.Error(ctx, err.Error())
		ctx.JSON(http.StatusBadRequest, responseDto.GetErrorResponseObj(err.Error()))
		return
	}

	data, err := app.EmployeeLogic.GetEmpoyee(ctx, request.EmployeeId)
	if err != nil {
		logger.Error(ctx, err.Error())
		ctx.JSON(http.StatusInternalServerError, responseDto.GetErrorResponseObj(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, responseDto.GetSuccessResponseObj(data))
}
