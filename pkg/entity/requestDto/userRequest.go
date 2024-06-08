package requestDto

import "crud/pkg/entity"

type CreateEmployeeRequest struct {
	EmployeeData entity.Employee `json:"employeeData" validate:"required"`
}

type UpdateEmployeeRequest struct {
	EmployeeData entity.Employee `json:"employeeData" validate:"required"`
}

type GetEmployeeRequest struct {
	EmployeeId string `json:"employeeId" validate:"required"`
}

type DeleteEmployeeRequest struct {
	EmployeeId string `json:"employeeId" validate:"required"`
}
