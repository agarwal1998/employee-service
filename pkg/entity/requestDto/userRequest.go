package requestDto

import "crud/pkg/entity"

type CreateEmployeeRequest struct {
	EmployeeData entity.Employee `json:"employeeData" validate:"required"`
}

type UpdateEmployeeRequest struct {
	EmployeeData entity.Employee `json:"employeeData" validate:"required"`
}

type GetEmployeeRequest struct {
	EmployeeId int `json:"employeeId" validate:"required"`
}

type DeleteEmployeeRequest struct {
	EmployeeId int `json:"employeeId" validate:"required"`
}
