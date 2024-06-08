package logic

import (
	"context"
	"crud/pkg/entity"
	repository "crud/pkg/repository/interface"
)

type EmployeeLogic struct {
	EmployeeRepo repository.UserRepoInterface
}

func NewEmployeeLogic(userRepo repository.UserRepoInterface) *EmployeeLogic {
	return &EmployeeLogic{userRepo}
}

func (el *EmployeeLogic) GetEmpoyee(ctx context.Context, employeeId int) (entity.Employee, error) {
	return el.EmployeeRepo.GetEmpoyee(ctx, employeeId)
}

func (el *EmployeeLogic) UpdateEmpoyee(ctx context.Context, updateReq entity.Employee) error {
	return el.EmployeeRepo.UpdateEmpoyee(ctx, updateReq)
}

func (el *EmployeeLogic) DeleteEmpoyee(ctx context.Context, employeeId int) error {
	return el.EmployeeRepo.DeleteEmpoyee(ctx, employeeId)

}

func (el *EmployeeLogic) Creatempoyee(ctx context.Context, employee entity.Employee) error {
	return el.EmployeeRepo.Creatempoyee(ctx, employee)
}
