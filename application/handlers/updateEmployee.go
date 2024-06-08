package handlers

import (
	"crud/common/logger"
	"crud/pkg/entity/requestDto"
	"crud/pkg/entity/responseDto"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateCreateUpdateRequest(ctx *gin.Context, req *requestDto.UpdateEmployeeRequest) error {
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		return err
	}
	if req.EmployeeData.ID == nil {
		return errors.New("Employee Id is required")
	}
	return nil
}

func (app *ApiHandler) UpdateEmpoyee(ctx *gin.Context) {
	var request requestDto.UpdateEmployeeRequest
	err := validateCreateUpdateRequest(ctx, &request)
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
