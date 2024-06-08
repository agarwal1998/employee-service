package application

import (
	"crud/application/handlers"
	cloudSql "crud/drivers"
	"crud/pkg/logic"
	"crud/pkg/repository/infrastructure"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	sqlClient := cloudSql.ConnectDB()
	userRepo := infrastructure.NewEmployeeRepo(sqlClient)
	employeeLogic := logic.NewEmployeeLogic(userRepo)
	apiHandler := handlers.NewApiHandler(employeeLogic)
	router := gin.Default()
	v1 := router.Group("/employee-service/v1")
	v1.GET("/employee/:employeeId", apiHandler.GetEmpoyee)
	v1.POST("/employee", apiHandler.CreateEmpoyee)
	v1.PUT("/employee", apiHandler.UpdateEmpoyee)
	v1.DELETE("/employee/:employeeId", apiHandler.DeleteEmpoyee)
	return router
}
