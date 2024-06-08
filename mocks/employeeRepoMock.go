package mocks

import (
	"context"
	"crud/pkg/entity"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepoMock struct {
	mock.Mock
}

func (r EmployeeRepoMock) GetEmpoyee(ctx context.Context, employeeId int) (entity.Employee, error) {
	args := r.Called(ctx, employeeId)
	return args.Get(0).(entity.Employee), nil
}
func (r EmployeeRepoMock) UpdateEmpoyee(ctx context.Context, updateReq entity.Employee) error {
	r.Called(ctx, updateReq)
	return nil
}
func (r EmployeeRepoMock) DeleteEmpoyee(ctx context.Context, employeeId int) error {
	r.Called(ctx, employeeId)
	return nil
}
func (r EmployeeRepoMock) Creatempoyee(ctx context.Context, employee entity.Employee) error {
	r.Called(ctx, employee)
	return nil
}
