package handlers

import (
	"crud/common/logger"
	"crud/pkg/entity/requestDto"
	"crud/pkg/entity/responseDto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *ApiHandler) DeleteEmpoyee(ctx *gin.Context) {
	var request requestDto.DeleteEmployeeRequest
	err := ctx.ShouldBindQuery(&request)
	if err != nil {
		logger.Error(ctx, err.Error())
		ctx.JSON(http.StatusBadRequest, responseDto.GetErrorResponseObj(err.Error()))
		return
	}

	err = app.EmployeeLogic.DeleteEmpoyee(ctx, request.EmployeeId)
	if err != nil {
		logger.Error(ctx, err.Error())
		ctx.JSON(http.StatusInternalServerError, responseDto.GetErrorResponseObj(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, responseDto.GetSuccessResponseObj(nil))
}
