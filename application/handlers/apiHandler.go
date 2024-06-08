package handlers

import "crud/pkg/logic"

type ApiHandler struct {
	EmployeeLogic *logic.EmployeeLogic
}

func NewApiHandler(EmployeeLogic *logic.EmployeeLogic) *ApiHandler {
	return &ApiHandler{EmployeeLogic}
}
