package handlers

import (
	"crud/common/logger"
	"crud/pkg/entity/requestDto"
	"crud/pkg/entity/responseDto"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateCreateEmpoyeeRequest(ctx *gin.Context, req *requestDto.CreateEmployeeRequest) error {
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		return err
	}
	if req.EmployeeData.Name == nil {
		return errors.New("Employee name is required")
	}
	return nil
}

func (app *ApiHandler) CreateEmpoyee(ctx *gin.Context) {
	var request requestDto.CreateEmployeeRequest
	err := validateCreateEmpoyeeRequest(ctx, &request)
	if err != nil {
		logger.Error(ctx, err.Error())
		ctx.JSON(http.StatusBadRequest, responseDto.GetErrorResponseObj(err.Error()))
		return
	}

	err = app.EmployeeLogic.Creatempoyee(ctx, request.EmployeeData)
	if err != nil {
		logger.Error(ctx, err.Error())
		ctx.JSON(http.StatusInternalServerError, responseDto.GetErrorResponseObj(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, responseDto.GetSuccessResponseObj(nil))
}
